// Author: yann
// Date: 2022/5/12
// Desc: ginwrap

package constants

const (
	CodeSuccess        = 100000 //成功
	CodeBadRequest     = 100400 //非法请求
	CodeAuthFailed     = 100401 //认证失败
	CodeGrantFailed    = 100403 //授权失败
	CodeServerError    = 100500 //服务异常
	CodeTypeAssertFail = 100501 //类型断言失败
)

const (
	MsgSuccess        = "成功"
	MsgBadRequest     = "非法请求"
	MsgAuthFailed     = "认证失败"
	MsgGrantFailed    = "授权失败"
	MsgServerError    = "服务异常"
	MsgTypeAssertFail = "类型断言失败"
)
