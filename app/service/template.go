package service

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

var funcMap = template.FuncMap{
	"add": func(i, j int) int {
		return i + j
	},

	"sub": func(i, j int) int {
		return i - j
	},

	"mul": func(i int, j float32) int {
		return int(float32(i) * j)
	},

	"eq": func(str1, str2 string) bool {
		return str1 == str2
	},

	"not": func(b bool) bool {
		return !b
	},

	"removeYear": func(date string) string {
		return date[:6]
	},

	"getPicSequence": func(seq, maxPostSeq int) int {
		result := seq - 13 // 13 is used here because the pics is 13 less than the posts
		if result < 0 {
			result = maxPostSeq - 13 + result
		}
		return result
	},

	"getPicPosition": func(seq, maxPostSeq int) string {
		if (maxPostSeq-seq)%2 == 1 {
			return "left"
		}
		return "right"
	},

	"getContentPosition": func(seq, maxPostSeq int) string {
		if (maxPostSeq-seq)%2 == 1 {
			return "right"
		}
		return "left"
	},
}

func (svc *Service) SetTemplates(server *gin.Engine) error {
	tmpl, err := template.New("").Funcs(funcMap).ParseGlob("./app/templates/*.tmpl")

	if err != nil {
		return err
	}
	server.SetHTMLTemplate(tmpl)
	return nil
}
