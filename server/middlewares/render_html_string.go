package middlewares

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"

	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"

	"github.com/MrHuxu/blogo/server/react"
)

func RenderHTMLString() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if strings.HasPrefix(ctx.Request.URL.String(), "/api") {
			return
		}

		pageInfo := getPageInfo(ctx)
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H(pageInfo))
	}
}

func getPageInfo(ctx *gin.Context) map[string]interface{} {
	status := ctx.Writer.Status()
	if status != 200 {
		urlVal, _ := react.VM.RunString(fmt.Sprintf("'/%d'", status))
		return map[string]interface{}{
			"meta":  fmt.Sprintf("%d Error", status),
			"title": "Error Happened!",
			"body":  template.HTML(react.Render(goja.FunctionCall{Arguments: []goja.Value{urlVal}}).String()),
		}
	}

	res, _ := ctx.Get("res")
	return convertResToPageInfo(ctx.Request.URL, res)
}

func convertResToPageInfo(url *url.URL, res interface{}) map[string]interface{} {
	var resMap map[string]interface{}
	if m, ok := res.(map[string]interface{}); ok {
		resMap = m
	}

	var meta, title, data interface{}
	if d, ok := resMap["meta"]; ok {
		meta = d
	}
	if d, ok := resMap["title"]; ok {
		title = d
	}
	if d, ok := resMap["data"]; ok {
		data = d
	}

	urlVal, _ := react.VM.RunString(fmt.Sprintf("'%s'", url))
	bytes, _ := json.Marshal(data)
	dataVal, _ := react.VM.RunString(fmt.Sprintf("'%s'", string(bytes)))
	bodyStr := react.Render(goja.FunctionCall{Arguments: []goja.Value{urlVal, dataVal}}).String()

	return map[string]interface{}{
		"meta":  meta,
		"title": title,
		"body":  template.HTML(bodyStr),
	}
}
