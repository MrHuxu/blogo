package server

import (
	"github.com/MrHuxu/blogo/server/handlers"
)

func (s *server) registerRoutes() {
	s.GET("/", handlers.DefaultPostHandler.SinglePage)
	s.GET("/api/page/:page", handlers.DefaultPostHandler.SinglePage)

	s.GET("/post/:title", handlers.DefaultPostHandler.SinglePost)

	s.GET("/tags", handlers.DefaultTagHandler.AllTags)
}
