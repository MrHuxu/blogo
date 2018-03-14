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

	"initYears": func() []string {
		return []string{}
	},

	"checkYearToRender": func(existingYears []string, year string) (newYears []string, needToRender bool) {
		for _, existingYear := range existingYears {
			if existingYear == year {
				newYears = existingYears
				needToRender = false
				return
			}
		}
		newYears = append(existingYears, year)
		needToRender = true
		return
	},

	"getPicSequence": func(seq, maxPostSeq int) int {
		if seq < 13 {
			return maxPostSeq - seq
		}
		return seq
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
