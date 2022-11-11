package ctx

type baseCtx struct {
	userId int
	ICtx
}

func NewBaseCtx(ctx ICtx, userId int) ICtx {
	return &baseCtx{
		userId: userId,
		ICtx:   ctx,
	}
}

func (ctx *baseCtx) UserID() int {
	return ctx.userId
}
