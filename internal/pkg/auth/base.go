// Author: yann
// Date: 2022/5/12
// Desc: auth

package auth

import (
	"fmt"
	"github.com/Hidata-xyz/go-example/internal/pkg/middleware"
	"github.com/Hidata-xyz/go-example/pkg/ginwrap"
	"github.com/Hidata-xyz/go-example/pkg/ginwrap/constants"
	"github.com/gin-gonic/gin"
)

type Base struct {
	UID string `json:"-"`
	GID string `json:"-"`
}

func GetBase(ctx *gin.Context) *Base {
	base := Base{
		UID: ctx.GetString(middleware.CtxUID),
	}
	if len(base.UID) > 3 {
		base.GID = base.UID[:3]
	}
	return &base
}

func MustGetBase(ctx *gin.Context) (*Base, error) {
	uid := ctx.GetString(middleware.CtxUID)
	if len(uid) < 3 {
		return nil, fmt.Errorf("UID未找到或者ID格式不正确: %s", uid)
	}
	return &Base{
		UID: uid,
		GID: uid[:3],
	}, nil
}

func BaseErrResponse(err error) *ginwrap.Response {
	return &ginwrap.Response{
		Code:    constants.CodeServerError,
		Message: err.Error(),
	}
}
