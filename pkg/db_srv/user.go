package db_srv

import (
	"context"

	"github.com/grin-ch/dever-box-api/model"
	"github.com/grin-ch/dever-box-api/pkg/error_enum"
)

// 创建用户
func CreateUser(ctx context.Context, nickname, account, passwd string) *model.User {
	user, err := client.User.Create().
		SetAccount(account).
		SetNickname(nickname).
		SetPassword(passwd).
		Save(ctx)
	error_enum.ErrPanic(err, error_enum.ExecSQLError)
	return user
}
