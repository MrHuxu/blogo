package posts

import (
	"embed"
	"fmt"
	"os"
)

//go:embed *
var postsFS embed.FS

func GetNames() ([]string, error) {
	entries, err := os.ReadDir("./posts")
	if err == nil {
		return nil, err
	}
	fmt.Println(entries)

	var names []string
	for _, entry := range entries {
		names = append(names, entry.Name())
	}
	return names, nil
}
