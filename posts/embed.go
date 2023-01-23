package posts

import (
	"embed"
)

//go:embed *
var postsFS embed.FS

// GetPosts ...
func GetPosts() ([]*Post, map[string]*Post, error) {
	entries, err := postsFS.ReadDir(".")
	if err != nil {
		return nil, nil, err
	}

	var posts []*Post
	mapTitleToPosts := make(map[string]*Post)
	for _, entry := range entries {
		if !ValidatePostFilename(entry.Name()) {
			continue
		}

		post := ConvFilenameToPost(entry.Name())
		posts = append(posts, post)
		mapTitleToPosts[post.Title] = post
	}

	return posts, mapTitleToPosts, nil
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
