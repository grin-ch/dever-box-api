package action

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/grin-ch/dever-box-api/api/web"
	"github.com/grin-ch/dever-box-api/auth"
	"github.com/grin-ch/dever-box-api/ctx"
	"github.com/grin-ch/dever-box-api/proxy"
)

var router *gin.Engine

func Router() *gin.Engine {
	router = gin.Default()
	unAuthApi()
	router.Use(auth.AuthMiddlewares()...)

	return router
}

func registery(ctx ctx.ICtx) {
	var url string
	if ctx.Module() != "" {
		url = fmt.Sprintf("/%s", ctx.Path())
	} else {
		url = fmt.Sprintf("/%s/%s", ctx.Module(), ctx.Path())
	}

	router.Handle(ctx.Method(), url, proxy.Around(ctx.Method(), ctx))
}

// 无需验证身份的接口
func unAuthApi() {
	registery(&web.Health{})
}
