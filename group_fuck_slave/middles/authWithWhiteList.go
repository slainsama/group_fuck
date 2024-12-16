package middles

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"group_fuck_slave/globals"
	"net/http"
)

// AuthWithWhiteList 验证请求是否来自白名单中的 IP 地址
func AuthWithWhiteList() gin.HandlerFunc {
	fmt.Println(globals.GlobalConfig.Secure.Whitelist)
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		println(clientIP)
		// 检查 IP 是否在白名单中
		if !isIPWhitelisted(clientIP) {
			// 如果不在白名单中，返回 404 状态码
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		// 如果在白名单中，继续处理请求
		c.Next()
	}
}

// 辅助函数：检查 IP 是否在白名单中
func isIPWhitelisted(ip string) bool {
	for _, whitelistedIP := range globals.GlobalConfig.Secure.Whitelist {
		println("here")
		println(whitelistedIP)
		if ip == whitelistedIP {
			return true
		}
	}
	return false
}
