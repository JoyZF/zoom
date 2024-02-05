// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/JoyZF/zoom/global"
)

// I118 add lang key to context.Context
func I118() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang := ctx.GetHeader(global.LangKey)
		if lang == "" {
			lang = global.ZH_CN
		}
		ctx.Set(global.LangKey, lang)
	}
}
