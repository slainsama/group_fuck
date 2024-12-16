package middles

import (
	"github.com/gin-gonic/gin"
	"group_fuck_slave/globals"
	"net/http"
)

func AuthWithKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Query("authKey")
		if key == globals.GlobalConfig.Secure.Key {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
	}
}
