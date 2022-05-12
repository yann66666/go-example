// Author: yann
// Date: 2022/5/11
// Desc: middleware

package middleware

import (
	"github.com/Hidata-xyz/go-example/pkg/ginwrap/constants"
	"github.com/gin-gonic/gin"
	"time"
)

var RecordTime = func(ctx *gin.Context) {
	ctx.Set(constants.CtxTime, time.Now().Nanosecond()/1e6)
	ctx.Next()
}
