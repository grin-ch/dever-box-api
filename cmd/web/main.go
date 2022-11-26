package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grin-ch/dever-box-api/cfg"
	"github.com/grin-ch/dever-box-api/cmd/web/action"
	"github.com/grin-ch/dever-box-api/util"
	"github.com/grin-ch/grin-utils/log"
	"github.com/grin-ch/grin-utils/tool"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	initCommon()

	gin.SetMode(cfg.Config.Server.Mode)
	r := action.Router()

	server := http.Server{
		Addr:    cfg.Config.Server.Addr,
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Logger.Errorf("server listen err:%s", err)
		}
	}()
	err := util.GracefulStop(server)
	if err != nil {
		log.Logger.Errorf("GracefulStop err:%s", err)
	}
}

// 初始化通用组件
func initCommon() {
	// 初始化日志
	log.InitLogger(
		log.WithPath(cfg.Config.Log.Path),
		log.WithLevel(cfg.Config.Log.Level),
		log.WithColor(cfg.Config.Log.HasCollor),
		log.WithCaller(cfg.Config.Log.HasCaller),
		log.WithMaxAge(time.Duration(cfg.Config.Log.MaxAge)*time.Second),
	)

	// 雪花算法
	tool.InitSnowflakeNode(cfg.Config.Server.Node)
}
