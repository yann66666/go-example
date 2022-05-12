// Author: yann
// Date: 2022/5/11
// Desc: ginwrap

package wrap

import (
	"fmt"
	"github.com/Hidata-xyz/go-example/pkg/ginwrap"
	"github.com/Hidata-xyz/go-example/pkg/ginwrap/constants"
	"github.com/gin-gonic/gin"
	"reflect"
)

type DecodeType uint8

const (
	DecodeTypeJson  DecodeType = iota + 1 //json
	DecodeTypeQuery                       //URL query参数
	DecodeTypeUri                         //URL path参数

)

type DecodeErr struct {
	Code int
	error
}

func NewDecodeErr(code int, error error) *DecodeErr {
	return &DecodeErr{Code: code, error: error}
}

func (d *DecodeErr) Response() *ginwrap.Response {
	var msg string
	if d.error != nil {
		msg = d.error.Error()
	}
	return &ginwrap.Response{
		Code:    d.Code,
		Message: msg,
	}
}

var (
	ErrNotPoint      = &DecodeErr{constants.CodeServerError, fmt.Errorf("不是指针")}
	ErrDecodeTypeErr = &DecodeErr{constants.CodeServerError, fmt.Errorf("不支持的解码类型")}
)

func (d DecodeType) Decode(ctx *gin.Context, dst interface{}) *DecodeErr {
	of := reflect.ValueOf(dst)
	if of.Kind() != reflect.Ptr {
		return ErrNotPoint
	}
	var err error
	switch d {
	case DecodeTypeJson:
		err = ctx.ShouldBindJSON(dst)
	case DecodeTypeQuery:
		err = ctx.ShouldBindQuery(dst)
	case DecodeTypeUri:
		err = ctx.ShouldBindUri(dst)
	default:
		return ErrDecodeTypeErr
	}

	if err == nil {
		return nil
	}

	err = GetError(err, dst)
	return NewDecodeErr(constants.CodeBadRequest, err)
}
