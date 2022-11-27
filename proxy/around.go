package proxy

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/grin-ch/dever-box-api/auth"
	"github.com/grin-ch/dever-box-api/cfg"
	"github.com/grin-ch/dever-box-api/ctx"
	"github.com/grin-ch/dever-box-api/error_enum"
	"github.com/grin-ch/grin-utils/log"
	"github.com/grin-ch/grin-utils/tool"
)

const (
	RequestID = "Requestid"
)

func Around(method string, ictx ctx.ICtx) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		reqID := fmt.Sprintf("%d", tool.NewSnowFlakeID())
		gctx.Header(RequestID, reqID)
		userId := 0
		value, has := gctx.Get(auth.UserKey)
		if has {
			userId = value.(auth.RoleBase).Id
		}
		baseCtx, cancel := ctx.NewBaseCtx(ictx, userId)
		defer cancel()
		defer func() {
			err := recover()
			if err != nil {
				baseCtx.ErrorHandle(err)
				gctx.Header("Content-Type", ctx.JSON)
				e, ok := err.(error_enum.ErrAble)
				if !ok {
					e = error_enum.UndefinedError(err)
				}
				gctx.JSON(200, gin.H{
					"code":    e.Code(),
					"msg":     e.Msg(),
					"err":     e.Err(),
					RequestID: reqID,
				})
			}
		}()
		baseCtx.Before()
		func() {
			cost := tool.Cost()
			gctx.Header("Content-Type", baseCtx.ContextType())
			defer func() {
				rsp := baseCtx.Action(gctx)
				deverLog(gctx, reqID, rsp)
				switch baseCtx.ContextType() {
				case ctx.STRING:
					gctx.String(200, fmt.Sprintf("%v", rsp))
				case ctx.STREAM: // 自由实现
				default:
					gctx.JSON(200, gin.H{
						"data":    rsp,
						"cost":    cost(),
						RequestID: reqID,
					})
				}
			}()
		}()
		baseCtx.After()
	}
}

func deverLog(gctx *gin.Context, reqID string, rsp any) {
	if cfg.Config.Server.Debug {
		log.Logger.Infof("%s:%sdata:%+v", RequestID, reqID, rsp)
	}
}
