// Author: yann
// Date: 2022/5/11
// Desc: middleware

package middleware

import (
	"fmt"
	"github.com/HiData-xyz/HiData.More/HiExtern/utils"
	"github.com/Hidata-xyz/go-example/pkg/ginwrap/constants"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"sync/atomic"
	"time"
)

var index uint64

func Print(c *gin.Context) {
	current := atomic.AddUint64(&index, 1)
	c.Next()
	end := time.Now()
	var f func(string) string
	if c.GetInt(constants.CtxCode) == constants.CodeSuccess {
		f = InfoColorMsg
	} else {
		f = ErrorColorMsg
	}
	fmt.Printf("[请求%d返回] %s | %d | %v | %v | %v | %s \n", current, end.Format(utils.TimeFormatSecond), c.Writer.Status(), end.UnixNano()/1e6-c.GetInt64(constants.CtxTime), f(cast.ToString(c.GetInt(constants.CtxCode))), f(c.GetString(constants.CtxMsg)), c.Request.URL.Path)
}

func InfoColorMsg(msg string) string {
	return fmt.Sprintf(" %c[%d;%d;%dm%s%c[0m ", 0x1B, 1, 0, 32, msg, 0x1B)
}

func ErrorColorMsg(msg string) string {
	return fmt.Sprintf(" %c[%d;%d;%dm%s%c[0m ", 0x1B, 1, 0, 31, msg, 0x1B)
}
