package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "This is blogo, xhu's blog in go.",
		})
	})
	r.Run("localhost:13109")
}
