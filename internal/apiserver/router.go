package apiserver

import (
	"github.com/JoyZF/errors"
	"github.com/JoyZF/zoom/internal/pkg/code"
	"github.com/JoyZF/zoom/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"net/http"
)

func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine) {
}

func installController(g *gin.Engine) *gin.Engine {
	// Middlewares.
	auto := newAutoAuth()
	g.NoRoute(auto.AuthFunc(), func(c *gin.Context) {
		core.WriteResponse(c, errors.WithCode(code.NotFound, code.GetMsg(code.NotFound, code.ZH_CN)), nil)
	})

	// v1 handlers, requiring authentication
	v1 := g.Group("/v1")
	{
		v1.GET("/index", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "init")
		})
	}

	return g
}

func newAutoAuth() middleware.AuthStrategy {
	return nil
	//return auth.NewAutoStrategy(newBasicAuth().(auth.BasicStrategy), newJWTAuth().(auth.JWTStrategy))
}
