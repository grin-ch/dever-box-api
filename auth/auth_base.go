package auth

import "github.com/gin-gonic/gin"

const (
	TokenKey = "token"
)

func AuthMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		authBase,
	}
}

func authBase(ctx *gin.Context) {
	token := ctx.GetHeader(TokenKey)
	if token == "" {
		ctx.Abort()
		return
	}

	_, err := ParseJWT(token)
	if err != nil {
		ctx.Abort()
		return
	}
}
