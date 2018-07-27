package react

import (
	"github.com/dop251/goja"
)

var (
	VM     *goja.Runtime
	Render func(goja.FunctionCall) goja.Value
)

func init() {
	initVM()
	initRenderer()
}
