package app

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"runtime"
	"strings"
)

type App struct {
	basename    string
	name        string
	description string
	options     CliOptions
	runFunc     RunFunc
	silence     bool
	noVersion   bool
	noConfig    bool
	//commands  []*Command
	args cobra.PositionalArgs
	cmd  *cobra.Command
}

// RunFunc defines the application's startup callback function.
type RunFunc func(basename string) error

type Option func(*App)

func NewApp(name, basename string, opts ...Option) *App {
	a := App{
		name:     name,
		basename: basename,
	}
	for _, o := range opts {
		o(&a)
	}
	// TODO buildCommand

	return &a
}

func WithOptions(opt CliOptions) Option {
	return func(a *App) {
		a.options = opt
	}
}

func WithDescription(commandDesc string) Option {
	return func(a *App) {
		a.description = commandDesc
	}
}

func WithRunFunc(runFunc RunFunc) Option {
	return func(a *App) {
		a.runFunc = runFunc
	}
}

func FormatBaseName(basename string) string {
	// Make case-insensitive and strip executable suffix if present
	if runtime.GOOS == "windows" {
		basename = strings.ToLower(basename)
		basename = strings.TrimSuffix(basename, ".exe")
	}

	return basename
}

func (a *App) Run() {
	if err := a.cmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}
