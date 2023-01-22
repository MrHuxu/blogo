package handlers

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/MrHuxu/blogo/server/conf"
)

var DefaultTagHandler TagHandler

type TagHandler interface {
	AllTags(*gin.Context)
}

func initTagHandler() {
	handler := &tagHandler{[]string{}, make(map[string]int)}
	handler.cacheTags()

	DefaultTagHandler = handler
}

type tagHandler struct {
	tags  []string
	times map[string]int
}

func (h *tagHandler) cacheTags() {
	entries, err := os.ReadDir(conf.Conf.Post.PostsPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), ".md") || strings.HasPrefix(entry.Name(), "WIP") {
			continue
		}

		for _, str := range strings.Split(strings.Split(strings.Split(entry.Name(), "#")[3], ".")[0], "-") {
			t := string(str)
			if _, ok := h.times[t]; ok {
				h.times[t]++
			} else {
				h.tags = append(h.tags, t)
				h.times[t] = 1
			}
		}
	}
}

func (h *tagHandler) AllTags(ctx *gin.Context) {
	res := make(map[string]interface{})
	res["meta"] = fmt.Sprintf("Life of xhu - Tags")
	res["title"] = fmt.Sprintf("Life of xhu - Tags")
	res["data"] = map[string]interface{}{
		"tags": map[string]interface{}{
			"tags":  h.tags,
			"times": h.times,
		},
	}
	ctx.Set("res", res)
}
