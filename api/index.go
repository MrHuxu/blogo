package api

import (
	"log"
	"net/http"
	"sort"

	"github.com/MrHuxu/blogo/posts"
	"github.com/MrHuxu/blogo/templates"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := templates.GetTemplate()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var data []*posts.Post
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			tmpl.Execute(w, map[string]any{
				"page":    "error",
				"title":   "Error",
				"message": err.Error(),
			})
			return
		}

		tmpl.Execute(w, map[string]any{
			"page":        "index",
			"title":       "Life of xhu",
			"selectedTag": r.URL.Query().Get("tag"),
			"posts":       data,
		})
	}()
	data, _, err = posts.GetPosts()
	if err != nil {
		log.Fatal(err)
	}
	if r.URL.Query().Has("tag") {
		data = filterPostsByTag(data, r.URL.Query().Get("tag"))
	}
	sort.Slice(data, func(i, j int) bool { return data[i].ID > data[j].ID })
}

func filterPostsByTag(original []*posts.Post, selectedTag string) []*posts.Post {
	var filtered []*posts.Post
	for _, post := range original {
		var match bool
		for _, tag := range post.Tags {
			if tag == selectedTag {
				match = true
				break
			}
		}
		if match {
			filtered = append(filtered, post)
		}
	}
	return filtered
}
