package cors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin") // 请求头部
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Origin", origin)
			// 允许的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			// 允许的请求头
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
			// 允许暴露的响应头
			c.Header("Access-Control-Expose-Headers", "Content-Length")
			// 预检请求的有效期，单位为秒
			c.Header("Access-Control-Max-Age", "600")
			// 是否允许发送Cookie
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		// 放行所有OPTIONS方法
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		// 处理请求
		c.Next()
	}
}
