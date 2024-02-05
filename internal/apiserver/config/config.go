// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import "github.com/JoyZF/zoom/internal/apiserver/options"

type Config struct {
	*options.Options `mapstructure:"options"`
}

func CreateConfigFromOptions(opts *options.Options) (*Config, error) {
	return &Config{opts}, nil
}
