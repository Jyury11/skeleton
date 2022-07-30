package vo

import (
	"github.com/pkg/errors"
)

// Error Constance
var (
	ErrInvalidInstanceName      = errors.New("invalid instance name")
	ErrInvalidInstanceContent   = errors.New("invalid instance content")
	ErrInvalidInstanceExtension = errors.New("invalid instance extension")
	ErrAlreadySetNextInstance   = errors.New("already set next instance")
)

// Instance Instance Value Object
type Instance struct {
	name      string
	content   string
	extension Extension
	children  []*Instance
}

// NewInstance Instance Value Object Constructor
func NewInstance(name, content string, extension Extension) (*Instance, error) {
	if name == "" {
		return nil, ErrInvalidInstanceName
	}
	if content == "" {
		return nil, ErrInvalidInstanceContent
	}
	if ExtensionNum <= extension {
		return nil, ErrInvalidInstanceExtension
	}
	t := &Instance{name: name, content: content, extension: extension}
	return t, nil
}

// NewEmptyContentInstance Instance Value Object Constructor
func NewEmptyContentInstance(name string, extension Extension) (*Instance, error) {
	if name == "" {
		return nil, ErrInvalidInstanceName
	}
	if ExtensionDirectory != extension {
		return nil, ErrInvalidInstanceExtension
	}
	t := &Instance{name: name, extension: extension}
	return t, nil
}

// Name Name Getter
func (t *Instance) Name() string {
	return t.name
}

// Content Content Getter
func (t *Instance) Content() string {
	return t.content
}

// Children Children Getter
func (t *Instance) Children() []*Instance {
	return t.children
}

// Extension Extension Getter
func (t *Instance) Extension() Extension {
	return t.extension
}

// SetChild Child Setter
func (t *Instance) SetChild(child *Instance) {
	if t.children == nil {
		t.children = []*Instance{
			child,
		}
		return
	}
	t.children = append(t.children, child)
}

// SetChildren Children Setter
func (t *Instance) SetChildren(children []*Instance) error {
	if t.children != nil {
		return ErrAlreadySetNextInstance
	}
	t.children = children
	return nil
}
