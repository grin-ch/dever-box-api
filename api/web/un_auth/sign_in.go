package un_auth

import (
	"github.com/gin-gonic/gin"
	"github.com/grin-ch/dever-box-api/pkg/ctx"
)

type SignIn struct {
	ctx.PostCtx
}

func (act *SignIn) Action(ctx *gin.Context) any {

	return gin.H{}
}

func (act *SignIn) Path() string {
	return "signin"
}
