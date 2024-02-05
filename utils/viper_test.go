// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"
	"testing"
)

type AppConfig struct {
	Database struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"database"`
}

func TestLoadConfig(t *testing.T) {
	v := LoadConfig()
	var config AppConfig
	err := v.Unmarshal(&config)
	if err != nil {
		fmt.Println("Unable to unmarshal config:", err)
		return
	}
	fmt.Printf("Database Host: %s\n", config.Database.Host)
	fmt.Printf("Database Port: %d\n", config.Database.Port)
}
