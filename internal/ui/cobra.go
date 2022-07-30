package ui

import (
	"context"
	"errors"

	"github.com/Jyury11/skeleton/internal/write/commands"
	"github.com/google/wire"
	"github.com/spf13/cobra"
)

var (
	ErrInvalidServiceName = errors.New("invalid service name")
	ErrInvalidSrc         = errors.New("invalid source path")
	ErrInvalidDst         = errors.New("invalid destination path")
)

// Cobra Cobra UI
type Cobra struct {
	rootCmd *cobra.Command
}

// NewCLI CLI Constructor
func NewCLI(rootCmd *cobra.Command) *Cobra {
	return &Cobra{rootCmd: rootCmd}
}

// Execute Execute
func (c *Cobra) Execute() error {
	return c.rootCmd.Execute()
}

// SetArgs Set args
func (c *Cobra) SetArgs(a []string) {
	c.rootCmd.SetArgs(a)
}

// CLISet Cli wire groupe
var CLISet = wire.NewSet(NewHandler, NewRootCmd, NewCLI)

// NewRootCmd Cobra Root Cmd Constructor
func NewRootCmd(h *Handler) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "skeleton",
		Short:   "cli",
		Long:    "cli",
		Version: "1.0",
	}
	cmd.AddCommand(h.Create())
	return cmd
}

// Handler Handler
type Handler struct {
	commands *commands.Commands
}

// NewHandler Handler Constructor
func NewHandler(commands *commands.Commands) *Handler {
	return &Handler{commands}
}

// Create Create Skeleton
func (h *Handler) Create() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create skeleton by template",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			s, err := cmd.Flags().GetString("service")
			if err != nil || s == "" {
				return ErrInvalidServiceName
			}

			src, err := cmd.Flags().GetString("src")
			if err != nil || src == "" {
				return ErrInvalidSrc
			}

			dst, err := cmd.Flags().GetString("dst")
			if err != nil || dst == "" {
				return ErrInvalidDst
			}

			v, err := cmd.Flags().GetString("values")
			if err != nil {
				return err
			}

			f, err := cmd.Flags().GetBool("force")
			if err != nil {
				return err
			}

			commands := commands.NewCreateCommand(s, src, dst, v, f)
			return h.commands.Create.Handle(ctx, commands)
		},
	}

	cmd.Flags().StringP("service", "s", "", "service name (required)")
	cmd.Flags().String("src", "", "source path (required)")
	cmd.Flags().String("dst", "", "destination path (required)")
	cmd.Flags().StringP("values", "v", "", "values yaml path")
	cmd.Flags().BoolP("force", "f", false, "always overwrite files")
	return cmd
}
