// Author: yann
// Date: 2022/5/11
// Desc: ginwrap

package ginwrap

import (
	"encoding/json"
	"fmt"
	"github.com/Hidata-xyz/go-example/pkg/ginwrap/constants"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Response 响应前端的固定格式
type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp int64       `json:"timestamp"` // 耗时  ms
}

const (
	RespParams = "响应[%s]: %s\n"
)

var PrintResp bool

func (r *Response) ReturnResult(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Headers", ctx.GetHeader("Access-Control-Request-Headers"))
	ctx.Header("Access-Control-Allow-Methods", ctx.GetHeader("Access-Control-Request-Method"))
	ctx.Header("Access-Control-Max-Age", "1728000")
	if PrintResp {
		indent, _ := json.Marshal(r)
		fmt.Printf(RespParams, ctx.Request.URL.Path, string(indent))
	}
	ctx.Set(constants.CtxCode, r.Code)
	ctx.Set(constants.CtxMsg, r.Message)
	r.Timestamp = time.Now().UnixNano()/1e6 - ctx.GetInt64(constants.CtxTime)
	ctx.AbortWithStatusJSON(http.StatusOK, r)
}

func Response400(err error) (response *Response) {
	response = new(Response)
	response.Code = constants.CodeBadRequest
	response.Message = err.Error()
	response.Timestamp = time.Now().UnixNano() / 1e6
	response.Data = struct{}{}
	return
}

func Response500(err error) (response *Response) {
	response = new(Response)
	response.Code = constants.CodeServerError
	response.Message = err.Error()
	response.Timestamp = time.Now().UnixNano() / 1e6
	response.Data = struct{}{}
	return
}

func Response200(data interface{}) (response *Response) {
	return &Response{
		Code:    constants.CodeSuccess,
		Message: constants.MsgSuccess,
		Data:    data,
	}
}
