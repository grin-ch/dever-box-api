package web

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grin-ch/dever-box-api/ctx"
)

type Health struct {
	ctx.GetCtx

	Grin string `req:"grin"`
}

func (act *Health) Action(ctx *gin.Context) any {
	now := time.Now()
	return gin.H{
		"time":    now.Format("2006-01-02 15:04:05"),
		"weekday": now.Weekday().String(),
		"userId":  act.UserID(),
	}
}

func (act *Health) Path() string {
	return "health"
}
