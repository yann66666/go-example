// Author: yann
// Date: 2022/5/11
// Desc: ginwrap

package wrap

import (
	"github.com/gin-gonic/gin"
)

type groupWrap struct {
	*gin.RouterGroup
}

func newGroupWrap(routerGroup *gin.RouterGroup) *groupWrap {
	return &groupWrap{RouterGroup: routerGroup}
}

func (g *groupWrap) Group(relativePath string) *groupWrap {
	return newGroupWrap(g.RouterGroup.Group(relativePath))
}

func (g *groupWrap) Handle(httpMethod, relativePath string, base IBase, handle Handle) {
	g.RouterGroup.Handle(httpMethod, relativePath, func(ctx *gin.Context) {
		var next IBase
		if base != nil {
			var err *DecodeErr
			if next, err = base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, next).ReturnResult(ctx)
	})
}

func (g *groupWrap) Any(relativePath string, base IBase, handle Handle) {
	g.RouterGroup.Any(relativePath, func(ctx *gin.Context) {
		var next IBase
		if base != nil {
			var err *DecodeErr
			if next, err = base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, next).ReturnResult(ctx)
	})
}

func (g *groupWrap) GET(relativePath string, base IBase, handle Handle) {
	g.RouterGroup.GET(relativePath, func(ctx *gin.Context) {
		var next IBase
		if base != nil {
			var err *DecodeErr
			if next, err = base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, next).ReturnResult(ctx)
	})
}

func (g *groupWrap) POST(relativePath string, base IBase, handle Handle) {
	g.RouterGroup.POST(relativePath, func(ctx *gin.Context) {
		var next IBase
		if base != nil {
			var err *DecodeErr
			if next, err = base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, next).ReturnResult(ctx)
	})
}

func (g *groupWrap) DELETE(relativePath string, base IBase, handle Handle) {
	g.RouterGroup.DELETE(relativePath, func(ctx *gin.Context) {
		var next IBase
		if base != nil {
			var err *DecodeErr
			if next, err = base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, next).ReturnResult(ctx)
	})
}

func (g *groupWrap) PATCH(relativePath string, base IBase, handle Handle) {
	g.RouterGroup.PATCH(relativePath, func(ctx *gin.Context) {
		var next IBase
		if base != nil {
			var err *DecodeErr
			if next, err = base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, next).ReturnResult(ctx)
	})
}

func (g *groupWrap) PUT(relativePath string, base IBase, handle Handle) {
	g.RouterGroup.PUT(relativePath, func(ctx *gin.Context) {
		var next IBase
		if base != nil {
			var err *DecodeErr
			if next, err = base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, next).ReturnResult(ctx)
	})
}

func (g *groupWrap) OPTIONS(relativePath string, base IBase, handle Handle) {
	g.RouterGroup.OPTIONS(relativePath, func(ctx *gin.Context) {
		var next IBase
		if base != nil {
			var err *DecodeErr
			if next, err = base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, next).ReturnResult(ctx)
	})
}

func (g *groupWrap) HEAD(relativePath string, base IBase, handle Handle) {
	g.RouterGroup.HEAD(relativePath, func(ctx *gin.Context) {
		var next IBase
		if base != nil {
			var err *DecodeErr
			if next, err = base.DecodeType().Decode(ctx, base); err != nil {
				err.Response().ReturnResult(ctx)
				return
			}
		}
		handle(ctx, next).ReturnResult(ctx)
	})
}
