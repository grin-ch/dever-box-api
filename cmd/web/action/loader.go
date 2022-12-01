package action

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grin-ch/dever-box-api/cfg"
	"github.com/grin-ch/dever-box-api/pkg/api/web/un_auth"
	"github.com/grin-ch/dever-box-api/pkg/api/web/user"
	"github.com/grin-ch/dever-box-api/pkg/ctx"
	"github.com/grin-ch/dever-box-api/pkg/middleware/auth"
	"github.com/grin-ch/dever-box-api/pkg/middleware/rate_limit"
	"github.com/grin-ch/dever-box-api/pkg/proxy"
)

var router *gin.Engine

func Router() *gin.Engine {
	router = gin.New()
	router.Use(cors())
	router.Use(rate_limit.RateLimiter(
		float64(cfg.Config.RateLimiter.Limit),
		cfg.Config.RateLimiter.Burst,
	)...)
	unAuthApi()
	router.Use(auth.AuthMiddlewares()...)

	registryUser()
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

func registery(act ctx.IAction) {
	var url string
	if act.Module() == "" {
		url = fmt.Sprintf("/%s", act.Path())
	} else {
		url = fmt.Sprintf("/%s/%s", act.Module(), act.Path())
	}

	router.Handle(act.Method(), url, proxy.Around(act.Method(), act))
}

// 无需验证身份的接口
func unAuthApi() {
	registery(&un_auth.Health{})
	registery(&un_auth.SignUp{})
	registery(&un_auth.SignIn{})
}

func registryUser() {
	registery((&user.Info{}))
}
