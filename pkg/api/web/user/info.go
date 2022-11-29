package user

import (
	"github.com/gin-gonic/gin"
	"github.com/grin-ch/dever-box-api/pkg/ctx"
)

type Info struct {
	ctx.GetCtx
}

func (act *Info) Action() any {
	return gin.H{
		"cliams": act.JwtCliam(),
	}
}

func (act *Info) Path() string {
	return "info"
}

func (act *Info) Module() string {
	return "user"
}
