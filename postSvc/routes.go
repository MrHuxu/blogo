package postSvc

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", ShowSnippets)
	server.GET("/page/*page", ShowSnippets)
	server.GET("/post/*title", ShowSinglePost)
}

func Global(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"prd":   "Production" == os.Getenv("ENV"),
		"title": "Blogo",
	})
}
