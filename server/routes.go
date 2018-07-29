package server

import (
	"github.com/MrHuxu/blogo/server/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *server) registerRoutes() {
	s.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/page/0")
	})

	s.GET("/page/:page", handlers.DefaultPostHandler.SinglePage)
	s.GET("/post/:title", handlers.DefaultPostHandler.SinglePost)

	s.GET("/tags", handlers.DefaultTagHandler.AllTags)
}
