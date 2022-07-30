package service

import (
	"github.com/Jyury11/skeleton/internal/write/domain/entity"
	"github.com/Jyury11/skeleton/internal/write/domain/vo"
)

// ConvertService Convert Domain Service
type ConvertService struct {
	buildService *BuildService
}

// NewConvertService Convert Domain Service Constructor
func NewConvertService(buildService *BuildService) *ConvertService {
	s := &ConvertService{buildService}
	return s
}

// Convert Convert Output Instance from Template
func (c *ConvertService) Convert(templates *entity.Templates, values map[string]interface{}) (*entity.Templates, error) {
	instances := make([]*vo.Instance, len(templates.Templates()))
	for i, temps := range templates.Templates() {
		ins, err := c.convert(temps, values)
		if err != nil {
			return nil, err
		}
		instances[i] = ins
	}
	templates.SetInstance(instances)
	return templates, nil
}

// convert Convert Output Instance from Template
func (c *ConvertService) convert(templates *vo.Template, values map[string]interface{}) (*vo.Instance, error) {
	ins, err := c.buildService.Build(templates, values)
	if err != nil {
		return nil, err
	}

	for _, temp := range templates.Children() {
		child, err := c.convert(temp, values)
		if err != nil {
			return nil, err
		}
		ins.SetChild(child)
	}
	return ins, nil
}
