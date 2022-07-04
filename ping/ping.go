package ping

import "github.com/gin-gonic/gin"

func PingRoute(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		message := "pong"
		value, exists := c.GetQuery("message")

		if exists {
			message = value
		}

		c.JSON(200, gin.H{
			"message": message,
		})
	})
}
