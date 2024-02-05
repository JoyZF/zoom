// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package options

import (
	"fmt"

	"github.com/JoyZF/zoom/internal/pkg/server"
)

// ServerRunOptions contains the options while running a generic api server.
type ServerRunOptions struct {
	Mode        string   `mapstructure:"gin-mode"`
	Healthz     bool     `mapstructure:"healthz"`
	Middlewares []string `mapstructure:"middlewares"`
	Port        int      `mapstructure:"port"`
	Address     string   `mapstructure:"bind-address"`
}

// NewServerRunOptions creates a new ServerRunOptions object with default parameters.
func NewServerRunOptions() *ServerRunOptions {
	return &ServerRunOptions{
		Mode:        "",
		Healthz:     false,
		Middlewares: nil,
	}
}

// ApplyTo applies the run options to the method receiver and returns self.
func (s *ServerRunOptions) ApplyTo(c *server.Config) error {
	c.Mode = s.Mode
	c.Healthz = s.Healthz
	c.Middlewares = s.Middlewares
	c.InsecureServing.Address = fmt.Sprintf("%s:%d", s.Address, s.Port)
	return nil
}

// Validate checks validation of ServerRunOptions.
func (s *ServerRunOptions) Validate() []error {
	errors := []error{}

	return errors
}
