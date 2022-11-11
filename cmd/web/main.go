package main

import (
	"github.com/gin-gonic/gin"
	"github.com/grin-ch/dever-box-api/cfg"
	"github.com/grin-ch/dever-box-api/cmd/web/action"
)

func main() {
	gin.SetMode(cfg.Config.Server.Mode)
	r := action.Router()
	r.Run(cfg.Config.Server.Addr)
}
