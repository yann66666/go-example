// Author: yann
// Date: 2022/5/11
// Desc: ginwrap

package wrap

import (
	"github.com/Hidata-xyz/go-example/pkg/ginwrap"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

type User struct {
	Name string `json:"name" form:"name" uri:"name" binding:"required" err:"用户名称不能为空"`
}

func (u User) DecodeType() DecodeType {
	return DecodeTypeQuery
}

func TestServer(t *testing.T) {
	engine := gin.New()
	wrap := NewEngineWrap(engine, EngineWrapLoggerOption(), PrintReqParamsOption(), PrintRespParamsOption())
	wrap.GET("", User{}, func(ctx *gin.Context, base IBase) *ginwrap.Response {
		return ginwrap.Response200(base)
	})
	http.ListenAndServe(":80", engine)
	select {}
}
