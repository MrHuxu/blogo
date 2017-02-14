package postSvc

import (
	"github.com/gin-gonic/gin"
)

func (pSvc *PostSvc) RegisterRoutes(server *gin.Engine) {
	server.GET("/", pSvc.ShowHomepage)
	server.GET("/post/:title", pSvc.ShowSinglePost)
}
