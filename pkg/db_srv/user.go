package db_srv

import (
	"context"

	"github.com/grin-ch/dever-box-api/model"
	"github.com/grin-ch/dever-box-api/model/user"
)

// CreateUser 创建用户
func CreateUser(ctx context.Context, nickname, account, passwd string) *model.User {
	user, err := client.User.Create().
		SetAccount(account).
		SetNickname(nickname).
		SetPassword(passwd).
		Save(ctx)
	return mustExec(user, err)
}

// FindUserByAccount 查询用户
func FindUserByAccount(ctx context.Context, account string) *model.User {
	user, err := client.User.Query().Where(user.Account(account)).Only(ctx)
	return mustExec(user, err)
}
