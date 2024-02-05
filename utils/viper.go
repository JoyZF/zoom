// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const (
	EnvPrefix = "ZOOM"
)

func loadEnv(v *viper.Viper, filePath string) error {
	// Open the .env file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	// Read each line and set environment variables
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			if err = os.Setenv(key, value); err != nil {
				panic("Error setting environment variable: " + key)
			}
			input := strings.Replace(strings.ToLower(strings.Replace(key, "ZOOM_", "", -1)), "_", ".", -1)
			fmt.Println(fmt.Sprintf("input: %s key: %s", input, key))
			_ = v.BindEnv(input, key)
		}
	}
	return scanner.Err()
}

func LoadConfig() *viper.Viper {
	v := viper.New()
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dirs := strings.Split(dir, "/")
	var temp string
	for _, d := range dirs {
		temp += d + "/"
		if _, err = os.Stat(temp + ".env"); os.IsNotExist(err) {
			continue
		} else {
			_ = loadEnv(v, temp+".env")
			break
		}
	}
	// zoom env prefix is ZOOM
	v.SetEnvPrefix(EnvPrefix)
	v.AutomaticEnv()
	return v
}
