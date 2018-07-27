package handlers

import (
	"github.com/gin-gonic/gin"
)

var DefaultTagHandler TagHandler

type TagHandler interface {
	AllTags(*gin.Context)
}

func initTagHandler() {
	handler := &tagHandler{}
	handler.cacheTags()

	DefaultTagHandler = handler
}

type tagHandler struct {
	tags  []tag
	times map[tag]int
}

type tag string

func (h *tagHandler) cacheTags() {

}

func (h *tagHandler) AllTags(*gin.Context) {}
