package api

import (
	"bytes"
	"fmt"
	"html/template"
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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var data *posts.Post
	defer func() {
		if err != nil {
			tmpl.Execute(w, map[string]any{
				"page":    "error",
				"title":   "Error",
				"message": err.Error(),
			})
			return
		}

		tmpl.Execute(w, map[string]any{
			"page":  "post",
			"title": fmt.Sprintf("Life of xhu - %s", data.Title),
			"post":  data,
		})
	}()
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		return
	}

	_, mapIDToPosts, err := posts.GetPosts()
	if err != nil {
		return
	}
	if data = mapIDToPosts[id]; data == nil {
		err = fmt.Errorf("Post[id=%d] not found", id)
		return
	}

	if err = data.LoadContent(); err != nil {
		return
	}
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(data.Content), &buf); err != nil {
		panic(err)
	}
	data.ContentHTML = template.HTML(buf.String())
}
