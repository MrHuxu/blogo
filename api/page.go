package api

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"

	"github.com/MrHuxu/blogo/config"
	"github.com/MrHuxu/blogo/posts"
	"github.com/MrHuxu/blogo/server/middlewares"
	"github.com/MrHuxu/blogo/server/templates"
)

// Page ...
func Page(w http.ResponseWriter, r *http.Request) {
	// template
	tmpl, err := templates.GetIndexTemplage()
	if err != nil {
		log.Fatal(err)
	}

	// config
	config, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	// page data
	var maxPage int
	var titles []string
	infos := make(map[string]*posts.Post)

	posts, _, err := posts.GetPosts()
	if err != nil {
		log.Fatal(err)
	}
	for _, post := range posts {
		titles = append(titles, post.Title)
		infos[post.Title] = post
		maxPage = int(math.Ceil(float64(len(titles)) / float64(config.Web.PerPage)))
	}
	sort.Slice(titles, func(i, j int) bool { return infos[titles[i]].Seq > infos[titles[j]].Seq })

	// url params
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page >= maxPage {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if page >= maxPage || page < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// render
	res := map[string]any{
		"meta":  fmt.Sprintf("Life of xhu - Page %d", page),
		"title": fmt.Sprintf("Life of xhu - Page %d", page),
		"data": map[string]any{
			"page": map[string]any{
				"titles":  getPagedTitles(titles, maxPage, page, config),
				"infos":   infos,
				"maxPage": maxPage,
			},
		},
	}
	pageInfo := middlewares.GetPageInfoFromRes(r.URL.String(), res)
	tmpl.Execute(w, pageInfo)
}

func getPagedTitles(titles []string, maxPage int, page int, config *config.Conf) []string {
	var pagedTitles []string
	if page == maxPage-1 {
		pagedTitles = titles[config.Web.PerPage*page:]
	} else {
		pagedTitles = titles[config.Web.PerPage*page : config.Web.PerPage*(page+1)]
	}
	return pagedTitles
}
