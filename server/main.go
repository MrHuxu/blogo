package main

import (
	"github.com/MrHuxu/blogo/server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/*path", routes.Global)
	server.LoadHTMLGlob("templates/*")

	const port = "13109"
	server.Run(":" + port)
}
