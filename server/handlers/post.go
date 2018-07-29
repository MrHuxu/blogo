package handlers

import (
	"fmt"
	"math"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/MrHuxu/blogo/server/conf"
)

var DefaultPostHandler PostHanler

type PostHanler interface {
	SinglePage(*gin.Context)
	SinglePost(*gin.Context)
}

func initPostHandler() {
	handler := &postHandler{0, []string{}, make(map[string]*post)}
	handler.cachePosts()

	DefaultPostHandler = handler
}

type postHandler struct {
	maxPage int
	titles  []string
	infos   map[string]*post
}

type post struct {
	Filename string    `json:"filename,omitempty"`
	Seq      int       `json:"seq,omitempty"`
	Title    string    `json:"title,omitempty"`
	Time     time.Time `json:"time,omitempty"`
	Tags     []tag     `json:"tags,omitempty"`
}

func (h *postHandler) cachePosts() {
	filepath.Walk(conf.Conf.Post.ArchivesPath, func(path string, _ os.FileInfo, _ error) error {
		tmp := strings.Split(path, "/")
		if len(tmp) > 1 && !strings.HasPrefix(tmp[1], "WIP") {
			p := convFilenameToPost(tmp[1])
			h.titles = append(h.titles, p.Title)
			h.infos[p.Title] = p
		}
		h.maxPage = int(math.Ceil(float64(len(h.titles)) / float64(conf.Conf.Web.PerPage)))
		return nil
	})
}

func (h *postHandler) SinglePage(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Param("page"))
	if err != nil || page >= h.maxPage {
		ctx.Status(http.StatusNotFound)
		return
	}

	res := make(map[string]interface{})
	res["meta"] = fmt.Sprintf("Life of xhu - Page %d", page)
	res["title"] = fmt.Sprintf("Life of xhu - Page %d", page)
	var pagedTitles []string
	if page == h.maxPage {
		pagedTitles = h.titles[conf.Conf.Web.PerPage*page:]
	} else {
		pagedTitles = h.titles[conf.Conf.Web.PerPage*page : conf.Conf.Web.PerPage*(page+1)-1]
	}

	res["data"] = map[string]interface{}{
		"page": map[string]interface{}{
			"titles": pagedTitles,
			"infos":  h.infos,
		},
	}

	ctx.Set("res", res)
}
func (h *postHandler) SinglePost(*gin.Context) {}

func convFilenameToPost(filename string) *post {
	p := &post{Filename: url.QueryEscape(filename), Tags: []tag{}}

	arr := strings.Split(filename, "*")
	if i, err := strconv.Atoi(strings.Trim(arr[0], "0")); err == nil {
		p.Seq = i
	}
	p.Title = url.QueryEscape(arr[1])
	if t, err := time.Parse("20060102", arr[2]); err == nil {
		p.Time = t
	}
	for _, str := range strings.Split(strings.Split(arr[3], ".")[0], "-") {
		p.Tags = append(p.Tags, tag(str))
	}

	return p
}
