package commands

import (
	"context"

	"github.com/friendsofgo/errors"
	"github.com/jyury11/skeleton/internal/write/domain/service"
	"github.com/jyury11/skeleton/internal/write/domain/vo"
	"github.com/jyury11/skeleton/internal/write/repository"
)

// CreateCommand CreateCommand
type CreateCommand struct {
	ServiceName string
	Src         string
	Dst         string
	Values      string
	IsForce     bool
}

// NewCreateCommand CreateCommand Constructor
func NewCreateCommand(
	serviceName string,
	src string,
	dst string,
	values string,
	isForce bool,
) CreateCommand {
	return CreateCommand{
		ServiceName: serviceName,
		Src:         src,
		Dst:         dst,
		Values:      values,
		IsForce:     isForce,
	}
}

// Create Create Interface
type Create interface {
	Handle(ctx context.Context, cmd CreateCommand) error
}

// create create
type create struct {
	repo           repository.Repository
	convertService *service.ConvertService
}

// NewCreate Create Constructor
func NewCreate(repo repository.Repository, convertService *service.ConvertService) *create {
	return &create{repo: repo, convertService: convertService}
}

// Handle Handle
func (h *create) Handle(ctx context.Context, cmd CreateCommand) error {
	o := vo.NewOption(cmd.IsForce)
	t, err := h.repo.Find(cmd.ServiceName, cmd.Src)
	if err != nil {
		return errors.Wrap(err, "find failed")
	}

	var v map[string]interface{}
	if cmd.Values != "" {
		v, err = h.repo.FindValues(cmd.Values)
		if err != nil {
			return errors.Wrap(err, "find values failed")
		}
	}
	v = t.MakeValues(v)

	t, err = h.convertService.Convert(t, v)
	if err != nil {
		return errors.Wrap(err, "build failed")
	}

	return h.repo.Save(cmd.Dst, t, o)
}
