package service

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/Jyury11/skeleton/internal/write/domain/vo"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	buildFunctionTable = map[string]interface{}{
		"Echo":              fmt.Println,
		"CamelToSnake":      wrapInterfaceToString(camelToSnake),
		"CamelToUpperSnake": wrapInterfaceToString(camelToUpperSnake),
		"SnakeToCamel":      wrapInterfaceToString(snakeToCamel),
		"ToUpper":           wrapInterfaceToString(strings.ToUpper),
		"ToLower":           wrapInterfaceToString(strings.ToLower),
		"ToTitle":           wrapInterfaceToString(strings.ToTitle),
		"Title":             wrapInterfaceToString(cases.Title(language.Und, cases.NoLower).String),
	}
)

// BuildService Build Domain Service
type BuildService struct {
}

// NewBuildService Build Domain Service Constructor
func NewBuildService() *BuildService {
	s := &BuildService{}
	return s
}

// Build build template
func (b *BuildService) Build(temp *vo.Template, values map[string]interface{}) (*vo.Instance, error) {
	name, err := b.build(temp.Name(), temp.Name(), values)
	if err != nil {
		return nil, err
	}

	if temp.Extension() == vo.ExtensionDirectory {
		return vo.NewEmptyContentInstance(name, vo.ExtensionDirectory)
	}

	content, err := b.build(name, temp.Content(), values)
	if err != nil {
		return nil, err
	}

	ins, err := vo.NewInstance(name, content, 1)
	if err != nil {
		return nil, err
	}
	return ins, nil
}

// build build template
func (b *BuildService) build(name, source string, values map[string]interface{}) (string, error) {
	tpl, err := template.New(name).Funcs(buildFunctionTable).Parse(source)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, values); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// wrapInterfaceToString interface convert string
func wrapInterfaceToString(callback func(string) string) func(interface{}) string {
	return func(any interface{}) string {
		switch t := any.(type) {
		case string:
			return callback(t)
		}
		return ""
	}
}

// camelToSnakeBase camel to snake case basic
func camelToSnakeBase(s string) string {
	if s == "" {
		return s
	}

	delimiter := "_"
	sLen := len(s)
	var snake string
	for i, current := range s {
		if i > 0 && i+1 < sLen {
			if current >= 'A' && current <= 'Z' {
				next := s[i+1]
				prev := s[i-1]
				if (next >= 'a' && next <= 'z') || (prev >= 'a' && prev <= 'z') {
					snake += delimiter
				}
			}
		}
		snake += string(current)
	}
	return snake
}

// camelToSnake camel to snake case
func camelToSnake(s string) string {
	return strings.ToLower(camelToSnakeBase(s))
}

// camelToUpperSnake camel to upper snake case
func camelToUpperSnake(s string) string {
	return strings.ToUpper(camelToSnakeBase(s))
}

// snakeToCamel snake to camel case
func snakeToCamel(s string) string {
	if s == "" {
		return s
	}

	delimiter := '_'
	sLen := len(s)
	var camel string
	for i, current := range s {
		if i < sLen {
			if current != delimiter {
				if i == 0 || rune(s[i-1]) == delimiter {
					camel += strings.ToUpper(string(current))
				} else {
					camel += string(current)
				}
			}
		}
	}

	return camel
}
