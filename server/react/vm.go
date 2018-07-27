package react

import (
	"github.com/dop251/goja"
)

// VM exports the instance of js runtime
var VM *goja.Runtime

func initVM() {
	VM = goja.New()
}
