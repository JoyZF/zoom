package options

import "github.com/JoyZF/zlog"

type LogOptions struct {
	OutputPath  string     `json:"output_path"        mapstructure:"output_path"`
	FileName    string     `json:"file_name"          mapstructure:"file_name"`
	Level       zlog.Level `json:"level"              mapstructure:"level"`
	Format      string     `json:"format"             mapstructure:"format"`
	ServiceName string     `json:"service_name"               mapstructure:"service_name"`
}

func NewLogOptions() *LogOptions {
	return &LogOptions{
		OutputPath:  "./logs/",
		FileName:    "app.log",
		Level:       zlog.DebugLevel,
		Format:      "json",
		ServiceName: "zoom",
	}
}

func NewZoomApiLogOptions() *LogOptions {
	return &LogOptions{
		OutputPath:  "./logs/",
		FileName:    "app.log",
		Level:       zlog.DebugLevel,
		Format:      "json",
		ServiceName: "zoom",
	}
}
