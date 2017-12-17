package main

import (
	"github.com/MrHuxu/blogo/app/service"
	"github.com/MrHuxu/blogo/app/util"
	"github.com/gin-gonic/gin"
)

const (
	assetDir   = "./app/assets"
	serverPort = "13109"
)

func main() {
	server := gin.Default()
	server.Static("/assets", assetDir)

	service, err := service.New()
	util.HandleError(err)

	err = service.SetTemplates(server)
	util.HandleError(err)

	service.RegisterRoutes(server)

	server.Run(":" + serverPort)
}
