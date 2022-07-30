//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/jyury11/skeleton/internal/ui"
	"github.com/jyury11/skeleton/internal/write/commands"
	"github.com/jyury11/skeleton/internal/write/domain/service"
	"github.com/jyury11/skeleton/internal/write/infra"
)

// InitializeLibrary Create DI Container
func InitializeLibrary() (*ui.Library, func(), error) {
	wire.Build(
		infra.NewFileRepository,
		service.NewBuildService,
		service.NewConvertService,
		commands.NewCommands,
		ui.LibSet,
	)
	return nil, nil, nil
}
