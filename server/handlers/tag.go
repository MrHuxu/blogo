package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/MrHuxu/blogo/server/conf"
)

var DefaultTagHandler TagHandler

type TagHandler interface {
	AllTags(*gin.Context)
}

func initTagHandler() {
	handler := &tagHandler{[]tag{}, make(map[tag]int)}
	handler.cacheTags()

	DefaultTagHandler = handler
}

type tagHandler struct {
	tags  []tag
	times map[tag]int
}

type tag string

func (h *tagHandler) cacheTags() {
	filepath.Walk(conf.Conf.Post.PostsPath, func(path string, _ os.FileInfo, _ error) error {
		tmp := strings.Split(path, "/")
		if len(tmp) > 1 && !strings.HasPrefix(tmp[1], "WIP") {
			for _, str := range strings.Split(strings.Split(strings.Split(tmp[1], "*")[3], ".")[0], "-") {
				t := tag(str)
				if _, ok := h.times[t]; ok {
					h.times[t]++
				} else {
					h.tags = append(h.tags, t)
					h.times[t] = 1
				}
			}
		}
		return nil
	})
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
