package main

import (
	"github.com/MrHuxu/blogo/app/service"
	"github.com/gin-gonic/gin"
)

const (
	assetDir   = "./app/assets"
	serverPort = "13109"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	server := gin.Default()
	server.Static("/assets", assetDir)

	service, err := service.New()
	handleError(err)

	err = service.SetTemplates(server)
	handleError(err)

	service.RegisterRoutes(server)

	server.Run(":" + serverPort)
}
