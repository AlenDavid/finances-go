package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
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
	r.Run() // listen and server on 0.0.0.0:8080
}
