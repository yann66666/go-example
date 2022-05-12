// Author: yann
// Date: 2022/5/11
// Desc: ginwrap

package wrap

import (
	"github.com/Hidata-xyz/go-example/pkg/ginwrap"
	"github.com/gin-gonic/gin"
)

type Handle func(ctx *gin.Context, params IBase) *ginwrap.Response
