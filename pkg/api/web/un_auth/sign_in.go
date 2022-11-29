package un_auth

import (
	"github.com/grin-ch/dever-box-api/pkg/cmp/user_cmp"
	"github.com/grin-ch/dever-box-api/pkg/ctx"
	"github.com/grin-ch/dever-box-api/pkg/error_enum"
)

type SignIn struct {
	ctx.PostCtx
}

func (act *SignIn) Action() any {
	account, passwd, has := act.GinCtx().Request.BasicAuth()
	error_enum.FalsePanic(has, error_enum.MissingParameter)
	return user_cmp.SignIn(act, account, passwd)
}

func (act *SignIn) Path() string {
	return "signin"
}
