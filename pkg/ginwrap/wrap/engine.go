// Author: yann
// Date: 2022/5/11
// Desc: ginwrap

package wrap

import (
	"github.com/gin-gonic/gin"
)

type EngineWrap struct {
	gin.IRouter
}

func NewEngineWrap(IRouter gin.IRouter) *EngineWrap {
	return &EngineWrap{IRouter: IRouter}
}

func (e *EngineWrap) Group(relativePath string) *GroupWrap {
	return NewGroupWrap(e.IRouter.Group(relativePath))
}

func (e *EngineWrap) Handle(httpMethod, relativePath string, base IBase, handle Handle) {
	e.IRouter.Handle(httpMethod, relativePath, func(ctx *gin.Context) {
		if err := base.DecodeType().Decode(ctx, base); err != nil {
			err.Response().ReturnResult(ctx)
			return
		}
		handle(ctx, base).ReturnResult(ctx)
	})
	return
}

func (e *EngineWrap) Any(relativePath string, base IBase, handle Handle) {
	e.IRouter.Any(relativePath, func(ctx *gin.Context) {
		if err := base.DecodeType().Decode(ctx, base); err != nil {
			err.Response().ReturnResult(ctx)
			return
		}
		handle(ctx, base).ReturnResult(ctx)
	})
}

func (e *EngineWrap) GET(relativePath string, base IBase, handle Handle) {
	e.IRouter.GET(relativePath, func(ctx *gin.Context) {
		if err := base.DecodeType().Decode(ctx, base); err != nil {
			err.Response().ReturnResult(ctx)
			return
		}
		handle(ctx, base).ReturnResult(ctx)
	})
}

func (e *EngineWrap) POST(relativePath string, base IBase, handle Handle) {
	e.IRouter.POST(relativePath, func(ctx *gin.Context) {
		if err := base.DecodeType().Decode(ctx, base); err != nil {
			err.Response().ReturnResult(ctx)
			return
		}
		handle(ctx, base).ReturnResult(ctx)
	})
}

func (e *EngineWrap) DELETE(relativePath string, base IBase, handle Handle) {
	e.IRouter.DELETE(relativePath, func(ctx *gin.Context) {
		if err := base.DecodeType().Decode(ctx, base); err != nil {
			err.Response().ReturnResult(ctx)
			return
		}
		handle(ctx, base).ReturnResult(ctx)
	})
}

func (e *EngineWrap) PATCH(relativePath string, base IBase, handle Handle) {
	e.IRouter.PATCH(relativePath, func(ctx *gin.Context) {
		if err := base.DecodeType().Decode(ctx, base); err != nil {
			err.Response().ReturnResult(ctx)
			return
		}
		handle(ctx, base).ReturnResult(ctx)
	})
}

func (e *EngineWrap) PUT(relativePath string, base IBase, handle Handle) {
	e.IRouter.PUT(relativePath, func(ctx *gin.Context) {
		if err := base.DecodeType().Decode(ctx, base); err != nil {
			err.Response().ReturnResult(ctx)
			return
		}
		handle(ctx, base).ReturnResult(ctx)
	})
}

func (e *EngineWrap) OPTIONS(relativePath string, base IBase, handle Handle) {
	e.IRouter.OPTIONS(relativePath, func(ctx *gin.Context) {
		if err := base.DecodeType().Decode(ctx, base); err != nil {
			err.Response().ReturnResult(ctx)
			return
		}
		handle(ctx, base).ReturnResult(ctx)
	})
}

func (e *EngineWrap) HEAD(relativePath string, base IBase, handle Handle) {
	e.IRouter.HEAD(relativePath, func(ctx *gin.Context) {
		if err := base.DecodeType().Decode(ctx, base); err != nil {
			err.Response().ReturnResult(ctx)
			return
		}
		handle(ctx, base).ReturnResult(ctx)
	})
}
