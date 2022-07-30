// Code generated by skeleton; DO NOT EDIT.

package usecase
import (
	"context"
	"log"

	"github.com/Jyury11/skeleton/example/dst/user/internal/entity"
	"github.com/Jyury11/skeleton/example/dst/user/internal/repository"
)

// CreateCommand ...
type CreateCommand struct {
	Id int
}

// NewCreateCommand ...
func NewCreateCommand(id int) CreateCommand {
	return CreateCommand{
		Id: id,
	}
}

// CreateUseCase ...
type CreateUseCase interface {
	Handle(ctx context.Context, cmd CreateCommand) error
}

// create ...
type create struct {
	repo repository.Repository
}

// NewCreate ...
func NewCreate(repo repository.Repository) *create {
	return &create{repo: repo}
}

// Handle ...
func (h *create) Handle(ctx context.Context, cmd CreateCommand) error {
	log.Println(entity.UserCreateEvent)
	p := entity.NewUser(cmd.Id)
	return h.repo.Save(p)
}