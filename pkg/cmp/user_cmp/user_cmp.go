package user_cmp

import (
	"github.com/grin-ch/dever-box-api/model"
	"github.com/grin-ch/dever-box-api/pkg/auth"
	"github.com/grin-ch/dever-box-api/pkg/ctx"
	"github.com/grin-ch/dever-box-api/pkg/db_srv"
	"github.com/grin-ch/dever-box-api/pkg/error_enum"
	"github.com/grin-ch/grin-utils/tool"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx ctx.ICtx, nickname, account, passwd string) *model.User {
	passwdbys, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	error_enum.ErrPanic(err, error_enum.UnknownError)
	user := db_srv.CreateUser(ctx, nickname, account, string(passwdbys))
	user.Password = ""
	return user
}

func SignIn(ctx ctx.ICtx, account, passwd string) string {
	user := db_srv.FindUserByAccount(ctx, account)
	error_enum.FalsePanic(user != nil, error_enum.AuthError)
	error_enum.ErrPanic(bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwd)), error_enum.AuthError)

	jwt, err := auth.GenerateJWT(auth.RoleBase{
		Id:       user.ID,
		UUID:     tool.MustUUIDv4(),
		Nickname: user.Nickname,
		Ip:       ctx.ClientIP(),
	})
	error_enum.ErrPanic(err, error_enum.UnknownError)
	return jwt
}
