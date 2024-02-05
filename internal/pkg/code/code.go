// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package code

import (
	"github.com/JoyZF/errors"

	"github.com/JoyZF/zoom/global"
)

// Base code
const (
	// Success
	// ZH_CN 成功
	// EN_US Success
	Success = 200
	// SystemError
	// ZH_CN 系统错误
	// EN_US SystemError
	SystemError = 500
	// NotFound
	// ZH_CN 未找到
	// EN_US NotFound
	NotFound = 404
)

const (
	ParamsError             = 10001
	GenericServiceErrorCode = 10002
)

func RegisterCoder() {
	for _, code := range CodeMap {
		if _, ok := code[global.ZH_CN]; ok {
			errors.Register(code[global.ZH_CN])
		}
		if _, ok := code[global.EN_US]; ok {
			errors.Register(code[global.EN_US])
		}
	}
}
