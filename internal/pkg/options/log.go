// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package options

import "github.com/JoyZF/zlog"

type LogOptions struct {
	OutputPath  string     `mapstructure:"output-path"`
	FileName    string     `mapstructure:"file-name"`
	Level       zlog.Level `mapstructure:"level"`
	Format      string     `mapstructure:"format"`
	ServiceName string     `mapstructure:"service-name"`
}

func NewLogOptions() *LogOptions {
	return &LogOptions{
		OutputPath:  "",
		FileName:    "",
		Level:       zlog.DebugLevel,
		Format:      "",
		ServiceName: "",
	}
}

func NewZoomApiLogOptions() *LogOptions {
	return &LogOptions{
		OutputPath:  "",
		FileName:    "",
		Level:       zlog.DebugLevel,
		Format:      "",
		ServiceName: "",
	}
}
