package apiserver

import (
	"github.com/JoyZF/zoom/internal/apiserver/config"
	"github.com/JoyZF/zoom/internal/apiserver/options"
	"github.com/JoyZF/zoom/internal/pkg/app"
	"github.com/JoyZF/zoom/internal/pkg/log"
)

const commandDesc = "zoom api server"

func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	app := app.NewApp("zoom api server",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithRunFunc(run(opts))) // FIXME WithRunFunc 很容易忽略掉
	return app
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		// init log
		log.Init(opts.LogOptions)

		// TODO init config
		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}

		return Run(cfg)
	}
}
