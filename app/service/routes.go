package service

import (
	"github.com/gin-gonic/gin"
)

func (svc *Service) RegisterRoutes(server *gin.Engine) {
	server.GET("/", svc.homeHandler)
	server.GET("/page/:page", svc.homeHandler)
	server.GET("/post/:title", svc.postHandler)
	server.GET("/archives", svc.archivesHandler)
}
