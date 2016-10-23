package main

import (
	"github.com/MrHuxu/blogo/postSvc"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")

	ps := postSvc.New()
	ps.RegisterRoutes(server)

	const port = "13109"
	server.Run(":" + port)
}
