package ui

import (
	"context"

	"github.com/google/wire"
	"github.com/jyury11/skeleton/internal/write/commands"
)

// Library Library UI
type Library struct {
	handler *HandlerLibrary
}

// NewCLI CLI Constructor
func NewLibrary(h *HandlerLibrary) *Library {
	return &Library{handler: h}
}

// Create Create
func (c *Library) Create(serviceName string, src string, dst string, values string, isForce bool) error {
	return c.handler.Create(serviceName, src, dst, values, isForce)
}

// LibSet Library wire groupe
var LibSet = wire.NewSet(NewHandlerLibrary, NewLibrary)

// HandlerLibrary HandlerLibrary
type HandlerLibrary struct {
	commands *commands.Commands
}

// NewHandlerLibrary HandlerLibrary Constructor
func NewHandlerLibrary(commands *commands.Commands) *HandlerLibrary {
	return &HandlerLibrary{commands}
}

// Create Create Skeleton
func (h *HandlerLibrary) Create(serviceName string, src string, dst string, values string, isForce bool) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	commands := commands.NewCreateCommand(serviceName, src, dst, values, isForce)
	return h.commands.Create.Handle(ctx, commands)
}
