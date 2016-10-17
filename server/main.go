package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/*path", Global)
	server.LoadHTMLGlob("templates/*")

	const port = "13109"
	server.Run(":" + port)
}
