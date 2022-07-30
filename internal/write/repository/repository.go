package repository

import (
	"github.com/jyury11/skeleton/internal/write/domain/entity"
	"github.com/jyury11/skeleton/internal/write/domain/vo"
)

// Repository Repository Interface
type Repository interface {
	Find(serviceName, src string) (*entity.Templates, error)
	FindValues(values string) (map[string]interface{}, error)
	Save(dst string, t *entity.Templates, option *vo.Option) error
}
