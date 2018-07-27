package server

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"

	"github.com/MrHuxu/blogo/server/conf"
	"github.com/MrHuxu/blogo/server/middlewares"
)

// DefaultServer exports an instance of interface Server
var DefaultServer Server

// Server defines the interface that server needs to implement
type Server interface {
	Run()
}

type server struct {
	*gin.Engine

	jsVM     *goja.Runtime
	jsRender func(goja.FunctionCall) goja.Value
}

func (s *server) Run() {
	s.Engine.Run(fmt.Sprintf(":%d", conf.Conf.Web.Port))
}

func init() {
	s := &server{
		Engine: gin.Default(),
	}
	s.Use(middlewares.RenderReactData())

	s.LoadHTMLGlob(conf.Conf.Web.TemplatesPath)
	s.registerRoutes()

	DefaultServer = s
}
