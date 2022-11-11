package auth

import "github.com/gin-gonic/gin"

const (
	TokenKey = "token"
	UserKey  = "RoleBase"
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

	cliams, err := ParseJWT(token)
	if err != nil {
		ctx.Abort()
		return
	}
	ctx.Set(UserKey, cliams.RoleBase)
	ctx.Next()
}
