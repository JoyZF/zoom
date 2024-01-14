package options

import "github.com/JoyZF/zoom/internal/pkg/options"

// Options runs a api server.
type Options struct {
	ServerRunOptions *options.ServerRunOptions
	GRPCOptions      *options.GRPCOptions
	MySQLOptions     *options.MySQLOptions
	RedisOptions     *options.RedisOptions
	JwtOptions       *options.JwtOptions
	LogOptions       *options.LogOptions
}

func NewOptions() *Options {
	return &Options{
		ServerRunOptions: options.NewServerRunOptions(),
		GRPCOptions:      options.NewGRPCOptions(),
		MySQLOptions:     options.NewMySQLOptions(),
		RedisOptions:     options.NewRedisOptions(),
		JwtOptions:       options.NewJwtOptions(),
		LogOptions:       options.NewZoomApiLogOptions(),
	}
}

func (o *Options) Validate() []error {
	return nil
}
