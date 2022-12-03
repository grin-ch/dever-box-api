package proxy

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grin-ch/dever-box-api/cfg"
	"github.com/grin-ch/dever-box-api/pkg/ctx"
	"github.com/grin-ch/dever-box-api/pkg/error_enum"
	"github.com/grin-ch/grin-utils/log"
	"github.com/grin-ch/grin-utils/tool"
)

const (
	RequestID = "Requestid"
)

func Around(method string, act ctx.IAction) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		reqID := gctx.GetHeader(RequestID)
		gctx.Header(RequestID, reqID)
		deadline := time.Duration(apiTimeOut(act.Module(), act.Path()))
		context, cancel := context.WithTimeout(gctx, deadline*time.Second)
		defer cancel()
		baseCtx := ctx.NewBaseCtx(context, gctx)
		baseCtx.Set(ctx.Method, act.Method())
		baseCtx.Set(ctx.Module, act.Module())
		baseCtx.Set(ctx.Path, act.Path())
		defer func() {
			err := recover()
			if err != nil {
				act.ErrorHandle(err)
				gctx.Header("Content-Type", ctx.JSON)
				e, ok := err.(error_enum.IErr)
				if !ok {
					e = error_enum.UndefinedError(err)
				}
				gctx.JSON(200, deverErr(e))
			}
		}()
		act.Before(baseCtx)
		error_enum.ErrPanic(gctx.ShouldBind(act), error_enum.MissingParameter)
		func() {
			cost := tool.Cost()
			gctx.Header("Content-Type", act.ContextType())
			defer func() {
				rsp := act.Action()
				deverLog(act, reqID, rsp)
				switch act.ContextType() {
				case ctx.STRING:
					gctx.String(200, fmt.Sprintf("%v", rsp))
				case ctx.JSON:
					gctx.JSON(200, gin.H{
						"data": rsp,
						"cost": cost(),
					})
				default: //自由处理
				}
			}()
		}()
		act.After(baseCtx)
	}
}

func apiTimeOut(module, path string) int {
	key := path
	if module != "" {
		key = module + "-" + key
	}
	if t, has := cfg.Config.ApiDeadline.Appoint[key]; has {
		return t
	}
	return cfg.Config.ApiDeadline.Default
}

func deverErr(e error_enum.IErr) gin.H {
	rsp := gin.H{
		"code": e.Code(),
		"msg":  e.Msg(),
	}
	if cfg.Config.Server.Debug {
		rsp["err"] = e.Err()
	}
	return rsp
}

func deverLog(act ctx.IAction, reqID string, rsp any) {
	if cfg.Config.Server.Debug {
		log.Logger.Infof("%s :%s IP:%s data:%+v", RequestID, reqID, act.ClientIP(), rsp)
	}
}
