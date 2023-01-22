package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"sort"
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

func (h *postHandler) cachePosts() {
	filepath.Walk(conf.Conf.Post.PostsPath, func(path string, _ os.FileInfo, _ error) error {
		tmp := strings.Split(path, "/")
		if len(tmp) > 1 && !strings.HasPrefix(tmp[1], "WIP") {
			p := convFilenameToPost(tmp[1])
			h.titles = append(h.titles, p.Title)
			h.infos[p.Title] = p
		}
		h.maxPage = int(math.Ceil(float64(len(h.titles)) / float64(conf.Conf.Web.PerPage)))
		return nil
	})
	sort.Slice(h.titles, func(i, j int) bool { return h.infos[h.titles[i]].Seq > h.infos[h.titles[j]].Seq })
}

func (h *postHandler) SinglePage(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Param("page"))
	if err != nil || page >= h.maxPage {
		ctx.Status(http.StatusNotFound)
		return
	}

	if page >= h.maxPage || page < 0 {
		ctx.Status(http.StatusNotFound)
		return
	}

	res := make(map[string]interface{})
	res["meta"] = fmt.Sprintf("Life of xhu - Page %d", page)
	res["title"] = fmt.Sprintf("Life of xhu - Page %d", page)
	res["data"] = map[string]interface{}{
		"page": map[string]interface{}{
			"titles":  h.getPagedTitles(page),
			"infos":   h.infos,
			"maxPage": h.maxPage,
		},
	}
	ctx.Set("res", res)
}

func (h *postHandler) getPagedTitles(page int) []string {
	var pagedTitles []string
	if page == h.maxPage-1 {
		pagedTitles = h.titles[conf.Conf.Web.PerPage*page:]
	} else {
		pagedTitles = h.titles[conf.Conf.Web.PerPage*page : conf.Conf.Web.PerPage*(page+1)]
	}
	return pagedTitles
}

func (h *postHandler) SinglePost(ctx *gin.Context) {
	title := ctx.Param("title")
	post, ok := h.infos[title]

	if !ok {
		ctx.Status(http.StatusNotFound)
		return
	}

	post.getContent()
	ctx.Set("res", map[string]interface{}{
		"meta":  fmt.Sprintf("Life of xhu - %s", post.Title),
		"title": fmt.Sprintf("Life of xhu - %s", post.Title),
		"data":  map[string]interface{}{"post": post},
	})
}

type post struct {
	Filename string    `json:"filename,omitempty"`
	Seq      int       `json:"seq,omitempty"`
	Title    string    `json:"title,omitempty"`
	Time     time.Time `json:"time,omitempty"`
	Tags     []tag     `json:"tags,omitempty"`
	Content  string    `json:"content,omitempty"`
}

func (p *post) getContent() {
	filename := p.Filename
	file, err := os.Open(fmt.Sprintf("%s/%s", conf.Conf.Post.PostsPath, filename))
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	p.Content = string(bytes)
}

func convFilenameToPost(filename string) *post {
	p := &post{Filename: filename, Tags: []tag{}}

	arr := strings.Split(filename, "*")
	if i, err := strconv.Atoi(strings.TrimLeft(arr[0], "0")); err == nil {
		p.Seq = i
	}
	p.Title = arr[1]
	if t, err := time.Parse("20060102", arr[2]); err == nil {
		p.Time = t
	}
	for _, str := range strings.Split(strings.Split(arr[3], ".")[0], "-") {
		p.Tags = append(p.Tags, tag(str))
	}

	return p
}
