package main

import (
	"fmt"
	"os"

	"github.com/MrHuxu/blogo/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.GET("/", func(c *gin.Context) {
		api.Index(c.Writer, c.Request)
	})
	r.GET("/post/:id", func(c *gin.Context) {
		c.Request.URL.RawQuery = fmt.Sprintf("id=%s", c.Param("id"))
		api.Post(c.Writer, c.Request)
	})
	r.Run(":" + os.Getenv("PORT"))
}
