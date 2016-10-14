package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Global(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"prd":   "Production" == os.Getenv("ENV"),
		"title": "Blogo",
	})
}
