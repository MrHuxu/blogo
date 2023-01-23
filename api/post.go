package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MrHuxu/blogo/posts"
	"github.com/MrHuxu/blogo/server/middlewares"
	"github.com/MrHuxu/blogo/server/templates"
)

// Post ...
func Post(w http.ResponseWriter, r *http.Request) {
	// template
	tmpl, err := templates.GetIndexTemplage()
	if err != nil {
		log.Fatal(err)
	}

	_, mapTitleToPosts, err := posts.GetPosts()
	if err != nil {
		log.Fatal(err)
	}

	post := mapTitleToPosts[r.URL.Query().Get("title")]
	if err = post.LoadContent(); err != nil {
		log.Fatal(err)
	}

	// render
	res := map[string]any{
		"meta":  fmt.Sprintf("Life of xhu - %s", post.Title),
		"title": fmt.Sprintf("Life of xhu - %s", post.Title),
		"data":  map[string]any{"post": post},
	}
	pageInfo := middlewares.GetPageInfoFromRes(r.URL.String(), res)
	tmpl.Execute(w, pageInfo)
}
