package un_auth

import (
	"github.com/grin-ch/dever-box-api/pkg/cmp/user_cmp"
	"github.com/grin-ch/dever-box-api/pkg/ctx"
)

type SignUp struct {
	ctx.PostCtx

	Nickname string `binding:"required"`
	Account  string `binding:"required"`
	Password string `binding:"required"`
}

func (act *SignUp) Action() any {
	return user_cmp.SignUp(act, act.Nickname, act.Account, act.Password)
}

func (act *SignUp) Path() string {
	return "signup"
}
