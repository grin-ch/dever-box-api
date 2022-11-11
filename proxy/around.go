package proxy

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/grin-ch/dever-box-api/auth"
	"github.com/grin-ch/dever-box-api/ctx"
	"github.com/grin-ch/grin-utils/tool"
)

func Around(method string, ictx ctx.ICtx) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		userId := 0
		value, has := gctx.Get(auth.UserKey)
		if has {
			userId = value.(auth.RoleBase).Id
		}
		baseCtx := ctx.NewBaseCtx(ictx, userId)
		defer func() {
			err := recover()
			if err != nil {
				ictx.ErrorHandle(err)
				gctx.Header("Content-Type", ctx.JSON)
				gctx.JSON(200, gin.H{
					"code": 400,
					"err":  fmt.Sprintf("%v", err),
				})
			}
		}()
		baseCtx.Before()
		func() {
			cost := tool.Cost()
			gctx.Header("Content-Type", baseCtx.ContextType())
			defer func() {
				rsp := baseCtx.Action(gctx)
				switch baseCtx.ContextType() {
				case ctx.STRING:
					gctx.String(200, fmt.Sprintf("%v", rsp))
				case ctx.STREAM: // 自由实现
				default:
					gctx.JSON(200, gin.H{
						"data": rsp,
						"cost": cost(),
					})
				}
			}()
		}()
		baseCtx.After()
	}
}
