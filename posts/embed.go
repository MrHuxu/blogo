package posts

import (
	"embed"
)

//go:embed *
var postsFS embed.FS

// GetPosts ...
func GetPosts() ([]*Post, map[int]*Post, error) {
	entries, err := postsFS.ReadDir(".")
	if err != nil {
		return nil, nil, err
	}

	var posts []*Post
	mapIDToPosts := make(map[int]*Post)
	for _, entry := range entries {
		if !ValidatePostFilename(entry.Name()) {
			continue
		}

		post := ConvFilenameToPost(entry.Name())
		posts = append(posts, post)
		mapIDToPosts[post.ID] = post
	}

	return posts, mapIDToPosts, nil
}

// LoadContent ...
func (p *Post) LoadContent() error {
	bytes, err := postsFS.ReadFile(p.Filename)
	if err != nil {
		return err
	}

	p.Content = string(bytes)
	return nil
}
