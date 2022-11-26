package ctx

import (
	"context"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grin-ch/grin-utils/log"
)

const (
	defaultStackSize = 4096
)

type baseCtx struct {
	context.Context
	ictx   ICtx
	userId int
}

func NewBaseCtx(ctx ICtx, userId int) (ICtx, context.CancelFunc) {
	c, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	return &baseCtx{
		Context: c,
		userId:  userId,
		ictx:    ctx,
	}, cancel
}

func (ctx *baseCtx) UserID() int {
	return ctx.userId
}

func (ctx *baseCtx) Before()                      { ctx.ictx.Before() }
func (ctx *baseCtx) Action(gctx *gin.Context) any { return ctx.ictx.Action(gctx) }
func (ctx *baseCtx) After()                       { ctx.ictx.After() }
func (ctx *baseCtx) ErrorHandle(err any) {
	var buf [defaultStackSize]byte
	n := runtime.Stack(buf[:], false)
	log.Logger.Errorf("Module:%s,Path:%s, Error: %v\nTrace: %s", ctx.ictx.Module(), ctx.ictx.Path(), err, buf[:n])
}
func (ctx *baseCtx) Method() string      { return ctx.ictx.Method() }
func (ctx *baseCtx) Module() string      { return ctx.ictx.Module() }
func (ctx *baseCtx) Path() string        { return ctx.ictx.Path() }
func (ctx *baseCtx) ContextType() string { return ctx.ictx.ContextType() }
