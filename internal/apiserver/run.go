// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import "github.com/JoyZF/zoom/internal/apiserver/config"

func Run(cfg *config.Config) error {
	server, err := createAPIServer(cfg)
	if err != nil {
		return err
	}

	return server.PrepareRun().Run()
}
