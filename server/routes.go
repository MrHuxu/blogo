package server

import (
	"net/http"

	"github.com/MrHuxu/blogo/server/handlers"
	"github.com/gin-gonic/gin"
)

func (s *server) registerRoutes() {
	s.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/page/0")
	})

	s.GET("/page/:page", handlers.DefaultPostHandler.SinglePage)
	s.GET("/post/:title", handlers.DefaultPostHandler.SinglePost)

	s.GET("/tags", handlers.DefaultTagHandler.AllTags)

	s.Static("/assets", "./server/assets")
}
