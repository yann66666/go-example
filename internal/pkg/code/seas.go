// Author: yann
// Date: 2022/5/12
// Desc: code

package code

import (
	"fmt"
	"github.com/Hidata-xyz/go-example/pkg/ginwrap"
)

// SeasBadRequest 请求参数非法
const (
	SeasBadRequest Code = 209900 //请求参数非法
)

// 模型相关
const (
	SaveModelFail         Code = 200101 //保存模型失败
	CannotCreateFuncModel Code = 200102 //不能创建功能模型
)

// 功能相关
const ()

var seasMsg = map[Code]string{
	SeasBadRequest:        "请求参数非法 ",
	SaveModelFail:         "保存模型失败 ",
	CannotCreateFuncModel: "不能创建功能模型 ",
}

func NewSeasErrResponse(code Code, err error) *ginwrap.Response {
	msg := seasMsg[code]
	if err != nil {
		msg = fmt.Sprintf("%s %s", seasMsg[code], err.Error())
	}
	return &ginwrap.Response{
		Code:    code.Int(),
		Message: msg,
	}
}
