// Author: yann
// Date: 2022/5/11
// Desc: middleware

package middleware

import (
	"github.com/gin-gonic/gin"
)

const (
	CtxUID = "uid"
)

func Auth(ctx *gin.Context) {
	//验证token得到用户id
	ctx.Set(CtxUID, "uwp_xx1")
	ctx.Next()
}
