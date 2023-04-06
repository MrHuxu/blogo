package templates

import (
	"embed"
	"html/template"
	"time"
)

//go:embed *
var templatesFS embed.FS

// GetTemplate ...
func GetTemplate() (*template.Template, error) {
	tmpl, err := template.New("base.tmpl").Funcs(template.FuncMap{
		"getYearOfTime":  getYearOfTime,
		"formatTime":     formatTime,
		"formatFullTime": formatFullTime,
	}).ParseFS(templatesFS, "base.tmpl", "index.tmpl", "post.tmpl", "tags.tmpl", "error.tmpl")
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func getYearOfTime(time time.Time) string {
	return time.Format("2006")
}

func formatTime(time time.Time) string {
	return time.Format("Jan 02")
}

func formatFullTime(time time.Time) string {
	return time.Format("Jan 02, 2006")
}
