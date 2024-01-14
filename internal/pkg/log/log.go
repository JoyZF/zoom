package log

import (
	"github.com/JoyZF/zlog"
	"github.com/JoyZF/zoom/internal/pkg/options"
	"strings"
)

func Init(options *options.LogOptions) {
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
