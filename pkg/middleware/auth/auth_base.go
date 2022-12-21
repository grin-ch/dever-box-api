package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/grin-ch/dever-box-api/pkg/error_enum"
)

const (
	TokenKey  = "Authorization"
	CliamsKey = "CliamsKey"
)

func AuthMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		authBase,
	}
}

func authBase(ctx *gin.Context) {
	token := ctx.GetHeader(TokenKey)
	if token != "" {
		token = strings.Replace(token, "Bearer ", "", 1)
		cliams, err := ParseJWT(token)
		if err == nil {
			ctx.Set(CliamsKey, cliams)
			ctx.Next()
			return
		}
	}

	ctx.JSON(http.StatusOK, error_enum.EnumData(error_enum.AuthError))
	ctx.Abort()
}
