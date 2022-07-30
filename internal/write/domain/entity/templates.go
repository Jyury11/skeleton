package entity

import (
	"errors"

	"github.com/jyury11/skeleton/internal/write/domain/vo"
)

// Error Constance
var (
	ErrInvalidServiceName = errors.New("invalid service name")
	ErrInvalidTemplates   = errors.New("invalid templates")
)

// Templates Templates Entity
type Templates struct {
	serviceName string
	templates   []*vo.Template
	instance    []*vo.Instance
}

// NewTemplates Templates Constructor
func NewTemplates(serviceName string, templates []*vo.Template) (*Templates, error) {
	if serviceName == "" {
		return nil, ErrInvalidServiceName
	}
	if templates == nil {
		return nil, ErrInvalidTemplates
	}
	t := &Templates{serviceName: serviceName, templates: templates}
	return t, nil
}

// Templates Templates Getter
func (t *Templates) Templates() []*vo.Template {
	return t.templates
}

// Instance Instance Getter
func (t *Templates) Instance() []*vo.Instance {
	return t.instance
}

// SetInstance Instance Setter
func (t *Templates) SetInstance(instances []*vo.Instance) {
	t.instance = instances
}

// MakeValues make values by args and templates fields
func (t *Templates) MakeValues(values map[string]interface{}) map[string]interface{} {
	var v map[string]interface{}
	if values == nil {
		v = map[string]interface{}{
			"ServiceName": t.serviceName,
		}
		return v
	}

	v = values
	values["ServiceName"] = t.serviceName
	return v
}
