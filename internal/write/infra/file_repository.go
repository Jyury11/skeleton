package infra

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/pkg/errors"

	"github.com/jyury11/skeleton/internal/write/domain/entity"
	"github.com/jyury11/skeleton/internal/write/domain/vo"
	"github.com/jyury11/skeleton/internal/write/repository"
)

const (
	logMsgAlreadyExists   = "%s already exists. skip it. If you want to overwrite it, use '-f' or go generate prefix"
	parentDirIndexPrefix  = 2
	secondExtensionPrefix = 2
)

// FileRepository File Repository
type FileRepository struct {
}

// FileRepository File Repository Constructor
func NewFileRepository() repository.Repository {
	r := &FileRepository{}
	return r
}

// Find Find Templates
func (r *FileRepository) Find(serviceName, src string) (*entity.Templates, error) {
	absolute, err := filepath.Abs(src)
	if err != nil {
		return nil, err
	}
	templates, err := r.findTemplates(absolute)
	if err != nil {
		return nil, err
	}
	return entity.NewTemplates(serviceName, templates)
}

// FindValues Find Values
func (r *FileRepository) FindValues(values string) (map[string]interface{}, error) {
	f, err := os.Open(values)
	if err != nil {
		return nil, err
	}
	var v map[string]interface{}
	if err := yaml.NewDecoder(f).Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}

// Save Save Instance
func (r *FileRepository) Save(dst string, t *entity.Templates, option *vo.Option) error {
	absolute, err := filepath.Abs(dst)
	if err != nil {
		return err
	}
	dirList := strings.Split(absolute, "/")
	for _, instances := range t.Instance() {
		err := r.save(dirList, instances, option)
		if err != nil {
			return err
		}
	}
	return nil
}

// save Save Instance
func (r *FileRepository) save(dirList []string, instance *vo.Instance, option *vo.Option) error {
	for _, child := range instance.Children() {
		err := r.save(append(dirList, child.Name()), child, option)
		if err != nil {
			return err
		}
	}

	dst := "/" + filepath.Join(dirList...)
	if instance.Extension() == vo.ExtensionDirectory {
		if r.exists(dst) {
			return nil
		}
		return os.MkdirAll(dst, os.ModePerm)
	}

	if r.exists(dst) {
		if !option.IsOverride(instance.Content()) {
			log.Printf(logMsgAlreadyExists, dst)
			return nil
		}

		pre, err := ioutil.ReadFile(dst)
		if err != nil {
			return err
		}
		if !option.IsOverride(string(pre)) {
			log.Printf(logMsgAlreadyExists, dst)
			return nil
		}
	}

	dstDir := filepath.Dir(dst)
	if !r.exists(dstDir) {
		err := os.MkdirAll(dstDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return ioutil.WriteFile(dst, []byte(instance.Content()), os.ModePerm)
}

// findTemplates Find Templates
func (r *FileRepository) findTemplates(src string) ([]*vo.Template, error) {
	var templates []*vo.Template
	srcPaths := strings.Split(src, "/")
	err := filepath.WalkDir(src, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return errors.Wrap(err, "failed filepath.Walk")
		}

		var t *vo.Template
		if info.IsDir() {
			t, err = r.makeDirectoryTemplate(path, info.Name())
		} else {
			t, err = r.makeFileTemplate(path, info.Name())
		}
		if err != nil {
			return err
		}

		paths := strings.Split(path, "/")
		cursors := templates
		for i, p := range paths {
			if i < len(srcPaths)-1 && srcPaths[i] == p {
				continue
			}

			if i == len(srcPaths)-1 && len(srcPaths) == len(paths) && info.IsDir() {
				templates = append(templates, t)
				break
			}

			for _, template := range cursors {
				if template.Name() == p {
					if i == len(paths)-parentDirIndexPrefix {
						template.SetChild(t)
						break
					}
					cursors = template.Children()
					continue
				}
			}
		}
		return nil
	})
	return templates, err
}

// makeFileTemplate Make Templates By File
func (r *FileRepository) makeFileTemplate(path, name string) (*vo.Template, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed read file")
	}
	extension := vo.NewExtension(r.getExtension(name))
	t, err := vo.NewTemplate(r.getName(name), string(content), extension)
	if err != nil {
		return nil, errors.Wrap(err, "failed create file template instance")
	}
	return t, nil
}

// makeDirectoryTemplate Make Templates By Directory
func (r *FileRepository) makeDirectoryTemplate(path, name string) (*vo.Template, error) {
	t, err := vo.NewEmptyContentTemplate(name, vo.ExtensionDirectory)
	if err != nil {
		return nil, errors.Wrap(err, "failed create directory template instance")
	}
	return t, nil
}

// exists already item enable/disable decision
func (r *FileRepository) exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// getExtension get extension by absolute file name
func (r *FileRepository) getExtension(name string) string {
	ext := filepath.Ext(name)
	if ext != ".tmpl" {
		return ext
	}
	ss := strings.Split(name, ".")
	return "." + ss[len(ss)-secondExtensionPrefix]
}

// getName get name by absolute file name
func (r *FileRepository) getName(name string) string {
	ext := filepath.Ext(name)
	if ext != ".tmpl" {
		return name
	}
	ss := strings.Split(name, ".")
	return strings.Join(ss[:len(ss)-1], ".")
}
