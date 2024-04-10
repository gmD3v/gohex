package health

import "github.com/gin-gonic/gin"

func CheckHealth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	}
}
