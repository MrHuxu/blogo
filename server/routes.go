package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *server) registerRoutes() {
	s.GET("/", func(ctx *gin.Context) {
		ctx.Set("res", map[string]interface{}{
			"data": map[string]interface{}{
				"home": map[string]interface{}{
					"titles": []string{"test title"},
					"infos":  map[string]interface{}{"test title": 1},
				},
			},
		})
	})

	s.GET("/test", func(ctx *gin.Context) {
		ctx.Set("res", map[string]interface{}{
			"data": map[string]interface{}{
				"test": map[string]interface{}{
					"list":  []string{"test"},
					"infos": map[string]int{"test": 2},
				},
			},
		})
	})

	s.GET("/api/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"test": "this is a test",
		})
	})
}
