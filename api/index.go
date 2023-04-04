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
		log.Fatal(err)
	}

	posts, _, err := posts.GetPosts()
	if err != nil {
		log.Fatal(err)
	}
	if r.URL.Query().Has("tag") {
		posts = filterPostsByTag(posts, r.URL.Query().Get("tag"))
	}
	sort.Slice(posts, func(i, j int) bool { return posts[i].ID > posts[j].ID })

	tmpl.Execute(w, map[string]any{
		"page":        "index",
		"title":       "Life of xhu",
		"selectedTag": r.URL.Query().Get("tag"),
		"posts":       posts,
	})
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
