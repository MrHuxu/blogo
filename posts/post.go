package posts

import (
	"html/template"
	"strconv"
	"strings"
	"time"
)

type Post struct {
	Filename    string        `json:"filename,omitempty"`
	ID          int           `json:"id,omitempty"`
	Title       string        `json:"title,omitempty"`
	Time        time.Time     `json:"time,omitempty"`
	Tags        []string      `json:"tags,omitempty"`
	Content     string        `json:"content,omitempty"`
	ContentHTML template.HTML `json:"content_html,omitempty"`
}

// ConvFilenameToPost ...
func ConvFilenameToPost(filename string) *Post {
	p := &Post{Filename: filename, Tags: []string{}}

	arr := strings.Split(filename, "#")
	if i, err := strconv.Atoi(strings.TrimLeft(arr[0], "0")); err == nil {
		p.ID = i
	}
	p.Title = fixTitle(arr[1])
	if t, err := time.Parse("20060102", arr[2]); err == nil {
		p.Time = t
	}
	for _, str := range strings.Split(strings.Split(arr[3], ".")[0], "-") {
		p.Tags = append(p.Tags, string(str))
	}

	return p
}

func fixTitle(original string) string {
	return strings.Replace(
		strings.Replace(
			strings.Replace(
				original,
				"[", `"`, -1,
			),
			"]", `"`, -1,
		),
		"~", ":", -1,
	)
}

// ValidatePostFilename ...
func ValidatePostFilename(filename string) bool {
	return strings.HasSuffix(filename, ".md") && !strings.HasPrefix(filename, "WIP")
}
