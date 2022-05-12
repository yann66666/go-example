// Author: yann
// Date: 2022/5/11
// Desc: internal

package internal

import (
	"github.com/Hidata-xyz/go-example/internal/pkg/middleware"
	"github.com/Hidata-xyz/go-example/internal/smodel"
	"github.com/Hidata-xyz/go-example/pkg/ginwrap/wrap"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	engine := wrap.NewEngineWrap(router)
	engine.Use(middleware.Cors())
	engine.Use(middleware.Auth)
	group := engine.Group("manager")
	smodel.InitRouter(group.Group("seas"))
}
