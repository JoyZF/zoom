package apiserver

import (
	"github.com/JoyZF/zoom/internal/apiserver/config"
	"github.com/JoyZF/zoom/internal/apiserver/options"
	"github.com/JoyZF/zoom/internal/pkg/app"
)

const commandDesc = "zoom api server"

func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	app := app.NewApp("zoom api server",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithRunFunc(run(opts)))
	return app
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		// TODO init log
		//log.Init(opts.Log)
		//defer log.Flush()

		// TODO init config
		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}

		return Run(cfg)
	}
}
