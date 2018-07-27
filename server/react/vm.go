package react

import (
	"github.com/dop251/goja"
)

func initVM() {
	VM = goja.New()
}
