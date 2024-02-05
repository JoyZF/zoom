// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package options

import "github.com/JoyZF/zoom/internal/pkg/options"

// Options runs a api server.
type Options struct {
	ServerRunOptions *options.ServerRunOptions `mapstructure:"serverrunoptions"`
	GRPCOptions      *options.GRPCOptions      `mapstructure:"grpcoptions"`
	LogOptions       *options.LogOptions       `mapstructure:"logoptions"`
	StoreOptions     *options.StoreOptions     `mapstructure:"storeoptions"`
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) Validate() []error {
	return nil
}
