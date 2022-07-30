package vo

import "errors"

// Error Constance
var (
	ErrInvalidTemplateName      = errors.New("invalid template name")
	ErrInvalidTemplateContent   = errors.New("invalid template content")
	ErrInvalidTemplateExtension = errors.New("invalid template extension")
	ErrAlreadySetNextTemplate   = errors.New("already set next template")
)

// Template Template Value Object
type Template struct {
	name      string
	content   string
	extension Extension
	children  []*Template
}

// NewTemplate Template Value Object Constructor
func NewTemplate(name, content string, extension Extension) (*Template, error) {
	if name == "" {
		return nil, ErrInvalidTemplateName
	}
	if content == "" {
		return nil, ErrInvalidTemplateContent
	}
	if ExtensionNum <= extension {
		return nil, ErrInvalidTemplateExtension
	}
	t := &Template{name: name, content: content, extension: extension}
	return t, nil
}

// NewEmptyContentTemplate Template Value Object Constructor
func NewEmptyContentTemplate(name string, extension Extension) (*Template, error) {
	if name == "" {
		return nil, ErrInvalidTemplateName
	}
	if ExtensionDirectory != extension {
		return nil, ErrInvalidTemplateExtension
	}
	t := &Template{name: name, extension: extension}
	return t, nil
}

// Name Name Getter
func (t *Template) Name() string {
	return t.name
}

// Content Content Getter
func (t *Template) Content() string {
	return t.content
}

// Children Children Getter
func (t *Template) Children() []*Template {
	return t.children
}

// Extension Extension Getter
func (t *Template) Extension() Extension {
	return t.extension
}

// SetChild Child Setter
func (t *Template) SetChild(child *Template) {
	if t.children == nil {
		t.children = []*Template{
			child,
		}
		return
	}
	t.children = append(t.children, child)
}

// SetChildren Children Setter
func (t *Template) SetChildren(children []*Template) error {
	if t.children != nil {
		return ErrAlreadySetNextTemplate
	}
	t.children = children
	return nil
}
