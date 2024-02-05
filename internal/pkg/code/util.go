// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package code

import (
	"context"

	"github.com/JoyZF/errors"

	"github.com/JoyZF/zoom/global"
)

// ErrorWithCode return error with code
func ErrorWithCode(ctx context.Context, code int) error {
	return errors.WithCode(code, GetMsg(code, getLangFromCtx(ctx)))
}

func getLangFromCtx(ctx context.Context) string {
	value := ctx.Value(global.LangKey)
	if value == nil {
		return global.ZH_CN
	}
	if l, ok := value.(string); ok {
		return l
	}
	return global.ZH_CN
}
