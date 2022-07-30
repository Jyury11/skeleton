//go:build wireinject
// +build wireinject

package di

import (
	"github.com/Jyury11/skeleton/internal/ui"
	"github.com/Jyury11/skeleton/internal/write/commands"
	"github.com/Jyury11/skeleton/internal/write/domain/service"
	"github.com/Jyury11/skeleton/internal/write/infra"
	"github.com/google/wire"
)

// InitializeCLI Create DI Container
func InitializeCLI() (*ui.Cobra, func(), error) {
	wire.Build(
		infra.NewFileRepository,
		service.NewBuildService,
		service.NewConvertService,
		commands.NewCommands,
		ui.CLISet,
	)
	return nil, nil, nil
}
