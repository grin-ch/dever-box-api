package ctx

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	JSON   = "application/json"
	STRING = "application/text"
	STREAM = "application/octet-stream"
)

type ICtx interface {
	context.Context
	UserID() int
	Before()
	Action(*gin.Context) any
	After()
	ErrorHandle(any)

	Method() string
	Module() string
	Path() string
	ContextType() string
}

type GetCtx struct{ nopCtx }

func (GetCtx) Method() string { return http.MethodGet }

type PutCtx struct{ nopCtx }

func (PutCtx) Method() string { return http.MethodGet }

type PostCtx struct{ nopCtx }

func (PostCtx) Method() string { return http.MethodGet }

type DelCtx struct{ nopCtx }

func (DelCtx) Method() string { return http.MethodGet }

// no op ctx
type nopCtx struct{ context.Context }

func (nopCtx) UserID() int             { return 0 }
func (nopCtx) Before()                 {}
func (nopCtx) Action(*gin.Context) any { return nil }
func (nopCtx) After()                  {}
func (ctx nopCtx) ErrorHandle(err any) {}
func (nopCtx) Method() string          { return "" }
func (nopCtx) Module() string          { return "" }
func (nopCtx) Path() string            { return "" }
func (nopCtx) ContextType() string     { return JSON }
