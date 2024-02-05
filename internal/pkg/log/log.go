// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package log

import (
	"strings"

	"github.com/JoyZF/zlog"

	"github.com/JoyZF/zoom/internal/pkg/options"
)

func Init(options *options.LogOptions) {
	if options == nil {
		zlog.Fatalf("log options is nil")
	}
	var formatter zlog.Formatter
	switch strings.ToLower(options.Format) {
	case "json":
		formatter = &zlog.JsonFormatter{}
	case "text":
		formatter = &zlog.TextFormatter{}
	default:
		formatter = &zlog.JsonFormatter{}
	}
	zlog.New(zlog.WithServiceName(options.ServiceName),
		zlog.WithLevel(options.Level),
		zlog.WithFormatter(formatter),
		zlog.WithCleaner(&zlog.Clean{}),
		zlog.WithOutputPath(options.OutputPath, options.FileName),
		zlog.WithTraceKey("trace_id"))
}
