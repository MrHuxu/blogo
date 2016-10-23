package postSvc

import (
	"github.com/gin-gonic/gin"
)

func (pSvc *PostSvc) RegisterRoutes(server *gin.Engine) {
	server.GET("/", pSvc.ShowSnippets)
	server.GET("/page/:page", pSvc.ShowSnippets)
	server.GET("/archives", pSvc.ShowArchives)
	server.GET("/post/:title", pSvc.ShowSinglePost)
}
