package action

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grin-ch/dever-box-api/api/web/un_auth"
	"github.com/grin-ch/dever-box-api/pkg/auth"
	"github.com/grin-ch/dever-box-api/pkg/ctx"
	"github.com/grin-ch/dever-box-api/pkg/proxy"
)

var router *gin.Engine

func Router() *gin.Engine {
	router = gin.New()
	router.Use(cors())
	unAuthApi()
	router.Use(auth.AuthMiddlewares()...)

	return router
}

func cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		origin := ctx.Request.Header.Get("Origin")
		if origin != "" {
			ctx.Header("Access-Control-Allow-Origin", "*") // Access-Control-Allow-Origin * 替换为指定的域名
			ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			ctx.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			ctx.JSON(http.StatusOK, "ok!")
		}
		ctx.Next()
	}
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
	registery(&un_auth.Health{})
	registery(&un_auth.SignUp{})
	registery(&un_auth.SignIn{})
}
