package un_auth

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/grin-ch/dever-box-api/pkg/ctx"
	"github.com/grin-ch/dever-box-api/pkg/db_srv"
	"github.com/grin-ch/dever-box-api/pkg/error_enum"
	"golang.org/x/crypto/bcrypt"
)

type SignUp struct {
	ctx.PostCtx
}

type userParam struct {
	Nickname string `binding:"required"`
	Account  string `binding:"required"`
	Password string `binding:"required"`
}

func (act *SignUp) Action(ctx *gin.Context) any {
	var u userParam
	error_enum.ErrPanic(ctx.ShouldBind(&u), error_enum.MissingParameter)

	passwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	error_enum.ErrPanic(err, error_enum.UnknownError)
	user := db_srv.CreateUser(context.Background(), u.Nickname, u.Account, string(passwd))
	user.Password = ""
	return user
}

func (act *SignUp) Path() string {
	return "signup"
}
