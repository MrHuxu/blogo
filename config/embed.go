package config

import (
	"embed"
	"encoding/json"
)

// Conf ...
type Conf struct {
	Web struct {
		Port          int    `json:"port,omitempty"`
		TemplatesPath string `json:"templates_path,omitempty"`
		PerPage       int    `json:"per_page,omitempty"`
	} `json:"web,omitempty"`
	Post struct {
		PostsPath string `json:"posts_path,omitempty"`
	} `json:"post,omitempty"`
}

//go:embed *
var configFS embed.FS

// GetConfig ...
func GetConfig() (*Conf, error) {
	bytes, err := configFS.ReadFile("server.json")
	if err != nil {
		return nil, err
	}

	var conf Conf
	if err = json.Unmarshal(bytes, &conf); err != nil {
		return nil, err
	}
	return &conf, err
}
