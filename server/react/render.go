package react

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/dop251/goja"
)

func initRenderer() {
	file, err := os.Open("client/public/built/bundle.js")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)

	v, err := VM.RunString(string(bytes))
	if err != nil {
		log.Fatal(err)
	}

	Render = v.Export().(map[string]interface{})["renderHtmlString"].(func(goja.FunctionCall) goja.Value)
}
