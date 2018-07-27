package handlers

import (
	"math"
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
	filename string
	seq      int
	title    string
	time     time.Time
	tags     []tag
}

func (h *postHandler) cachePosts() {
	filepath.Walk(conf.Conf.Post.ArchivesPath, func(path string, _ os.FileInfo, _ error) error {
		tmp := strings.Split(path, "/")
		if len(tmp) > 1 && !strings.HasPrefix(tmp[1], "WIP") {
			p := convFilenameToPost(tmp[1])
			h.titles = append(h.titles, p.title)
			h.infos[p.title] = p
		}
		h.maxPage = int(math.Ceil(float64(len(h.titles)) / float64(conf.Conf.Web.PerPage)))
		return nil
	})
}

func (h *postHandler) SinglePage(*gin.Context) {}
func (h *postHandler) SinglePost(*gin.Context) {}

func convFilenameToPost(filename string) *post {
	p := &post{filename: filename, tags: []tag{}}

	arr := strings.Split(filename, "*")
	if i, err := strconv.Atoi(arr[0]); err != nil {
		p.seq = i
	}
	p.title = arr[1]
	if t, err := time.Parse("20060102", arr[2]); err != nil {
		p.time = t
	}
	for _, str := range strings.Split(strings.Split(arr[3], ".")[0], "-") {
		p.tags = append(p.tags, tag(str))
	}

	return p
}
