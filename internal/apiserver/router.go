// Copyright 2024 Joy <joyssss94@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"github.com/JoyZF/zoom/internal/apiserver/controller/v1/store"
	"github.com/gin-gonic/gin"

	"github.com/JoyZF/zoom/internal/pkg/code"
	"github.com/JoyZF/zoom/internal/pkg/middleware"
	"github.com/JoyZF/zoom/internal/pkg/response"
)

func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

// installMiddleware 额外注册中间件入口
func installMiddleware(g *gin.Engine) {

}

func installController(g *gin.Engine) *gin.Engine {
	g.NoRoute(func(ctx *gin.Context) {
		response.WriteResponse(ctx, code.ErrorWithCode(ctx, code.NotFound), nil)
	})

	// v1 handlers, requiring authentication
	v1 := g.Group("/v1")
	{
		sc := store.NewStoreController()
		v1.GET("/store", sc.Get)
		v1.PUT("/store", sc.Put)
		v1.PUT("/store/ttl", sc.PutWithTTL)
		v1.DELETE("/store", sc.Delete)
		v1.GET("/store/ttl", sc.TTL)
		v1.GET("/store/sync", sc.Sync)
		v1.GET("/store/stat", sc.Stat)
		v1.GET("/store/exist", sc.Exist)
		v1.GET("/store/expire", sc.Expire)
	}

	return g
}

func newAutoAuth() middleware.AuthStrategy {
	return nil
	//return auth.NewAutoStrategy(newBasicAuth().(auth.BasicStrategy), newJWTAuth().(auth.JWTStrategy))
}
