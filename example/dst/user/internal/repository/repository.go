// Code generated by skeleton; DO NOT EDIT.

package repository

import (
	"github.com/Jyury11/skeleton/example/dst/user/internal/entity"
)

// Repository ...
type Repository interface {
	Save(*entity.User) error
}
