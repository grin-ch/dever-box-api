package db_srv

import (
	"context"
	"sync"

	"github.com/grin-ch/dever-box-api/model"
	"github.com/grin-ch/dever-box-api/pkg/error_enum"
)

var (
	// 全局唯一的连接
	client   *model.Client
	initOnce sync.Once
)

// InitDB 初始化数据库
func InitDB(dsn string) func() error {
	initOnce.Do(func() {
		var err error
		client, err = model.Open("mysql", dsn)
		error_enum.ErrPanic(err, error_enum.ExecSQLError)
		error_enum.ErrPanic(client.Schema.Create(context.Background()), error_enum.ExecSQLError)
	})
	if client == nil {
		panic("InitDB error")
	}
	return client.Close
}
