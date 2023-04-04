package api

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/MrHuxu/blogo/posts"
	"github.com/MrHuxu/blogo/templates"
	"github.com/yuin/goldmark"
)

// Post ...
func Post(w http.ResponseWriter, r *http.Request) {
	tmpl, err := templates.GetTemplate()
	if err != nil {
		log.Fatal(err)
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	_, mapIDToPosts, err := posts.GetPosts()
	if err != nil {
		log.Fatal(err)
	}
	post := mapIDToPosts[id]

	post.LoadContent()
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(post.Content), &buf); err != nil {
		panic(err)
	}
	post.ContentHTML = template.HTML(buf.String())

	tmpl.Execute(w, map[string]any{
		"page": "post",
		"post": post,
	})
}
