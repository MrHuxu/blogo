package handlers

import (
	"github.com/gin-gonic/gin"
	"time"
)

var DefaultPostHandler PostHanler

type PostHanler interface {
	SinglePage(*gin.Context)
	SinglePost(*gin.Context)
}

func initPostHandler() {
	handler := &postHandler{}
	handler.cachePosts()

	DefaultPostHandler = handler
}

type postHandler struct {
	maxPage int
	titles  []string
	infos   map[string]*post
}

type post struct {
	fileName string
	seq      int
	title    string
	time     time.Time
	tags     []tag
}

func (h *postHandler) cachePosts() {

}

func (h *postHandler) SinglePage(*gin.Context) {}
func (h *postHandler) SinglePost(*gin.Context) {}
