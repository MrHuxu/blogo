package react

import (
	"embed"
	"encoding/json"
	"log"

	"github.com/dop251/goja"
)

//go:embed bundle.js
var bundleJsFS embed.FS

var jsRender func(goja.FunctionCall) goja.Value

func initRenderer() {
	bytes, err := bundleJsFS.ReadFile("bundle.js")
	if err != nil {
		log.Fatal(err)
	}

	v, err := VM.RunString(string(bytes))
	if err != nil {
		log.Fatal(err)
	}

	jsRender = v.Export().(map[string]interface{})["renderHtmlString"].(func(goja.FunctionCall) goja.Value)
}

// Render renders the react data to html string
func Render(url string, data interface{}) string {
	var m map[string]interface{}
	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, &m)
	return jsRender(goja.FunctionCall{Arguments: []goja.Value{VM.ToValue(url), VM.ToValue(m)}}).String()
}
