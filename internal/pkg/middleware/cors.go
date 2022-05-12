// Author: yann
// Date: 2022/5/12
// Desc: middleware

package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Cors() gin.HandlerFunc {
	// 允许所有跨域请求
	return func(c *gin.Context) {
		l := strings.Split(c.Request.URL.Path, "/")
		var NotAdd bool
		if len(l) > 2 {
			switch l[2] {
			case "gateway", "dockercontrol":
				NotAdd = true
			}
		}

		if !NotAdd {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		}
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
	}
}
