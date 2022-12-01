package ctx

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/grin-ch/dever-box-api/pkg/middleware/auth"
)

const (
	JSON   = "application/json"
	STRING = "application/text"
)

type ICtx interface {
	ICache[string, any]
	context.Context
	GinCtx() *gin.Context
	JwtCliam() *auth.Cliams
	ClientIP() string
}

func NewBaseCtx(ctx context.Context, gctx *gin.Context) ICtx {
	var c *auth.Cliams
	cliam, has := gctx.Get(auth.CliamsKey)
	if has {
		c = cliam.(*auth.Cliams)
	}
	bctx := &baseCtx{
		Context:  ctx,
		ctxCache: newCtxCache[string, any](),
		gctx:     gctx,
		cliam:    c,
	}
	return bctx
}

type baseCtx struct {
	context.Context
	*ctxCache[string, any]
	gctx  *gin.Context
	cliam *auth.Cliams
}

func (ctx *baseCtx) GinCtx() *gin.Context   { return ctx.gctx }
func (ctx *baseCtx) JwtCliam() *auth.Cliams { return ctx.cliam }
func (ctx *baseCtx) ClientIP() string       { return ctx.gctx.ClientIP() }
