package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	// r.Run("0.0.0.0:8080")
	r.RunTLS(
		"0.0.0.0:8080",
		"./envoy/certs/gin.example.com.crt",
		"./envoy/certs/gin.example.com.key",
	)
}
