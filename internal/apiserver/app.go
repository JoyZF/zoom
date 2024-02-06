// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"github.com/JoyZF/zlog"
	"github.com/JoyZF/zoom/pkg/store"

	"github.com/JoyZF/zoom/internal/apiserver/config"
	"github.com/JoyZF/zoom/internal/apiserver/options"
	"github.com/JoyZF/zoom/internal/pkg/app"
	"github.com/JoyZF/zoom/internal/pkg/code"
	"github.com/JoyZF/zoom/internal/pkg/log"
	"github.com/JoyZF/zoom/utils"
)

const commandDesc = "zoom api server"

func init() {
	code.RegisterCoder()
}

func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	// TODO register cronjob to sync data
	return app.NewApp("zoom api server",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithRunFunc(run(opts))) // FIXME WithRunFunc 很容易忽略掉
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		viper := utils.LoadConfig()
		var (
			cfg *config.Config
			err error
		)

		if cfg, err = config.CreateConfigFromOptions(opts); err != nil {
			return err
		}
		if err = viper.Unmarshal(cfg); err != nil {
			zlog.Fatalf("%v", err)
		}
		log.Init(cfg.LogOptions)
		if err = store.DB(cfg.StoreOptions.Driver); err != nil {
			zlog.Fatalf("init db driver err %+v", err)
		}
		return Run(cfg)
	}
}
