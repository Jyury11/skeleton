// Code generated by skeleton; DO NOT EDIT.

package infra

import (
	"fmt"

	"github.com/jyury11/skeleton/example/dst/user/internal/entity"
	"github.com/jyury11/skeleton/example/dst/user/internal/repository"
)

// StdRepository ...
type StdRepository struct {
}

// NewStdRepository ...
func NewStdRepository() repository.Repository {
	r := &StdRepository{}
	return r
}

// Save ...
func (s *StdRepository) Save(u *entity.User) error {
	fmt.Printf("user_id: %d\n", u.Id())
	return nil
}
