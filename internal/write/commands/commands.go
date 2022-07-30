package commands

import (
	"github.com/jyury11/skeleton/internal/write/domain/service"
	"github.com/jyury11/skeleton/internal/write/repository"
)

// Commands Commands
type Commands struct {
	Create
}

// NewCommands Commands Constructor
func NewCommands(repo repository.Repository, convertService *service.ConvertService) *Commands {
	return &Commands{
		Create: NewCreate(repo, convertService),
	}
}
