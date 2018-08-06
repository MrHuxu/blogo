package react

import (
	"strings"
)

func escape(str string, convFuncs ...convFunc) string {
	for _, fn := range convFuncs {
		str = fn(str)
	}
	return str
}

type convFunc func(string) string

func escapeSingleQuotes(str string) string {
	return strings.Replace(str, `'`, `\'`, -1)
}

func escapeDoubleQuotes(str string) string {
	return strings.Replace(str, `\"`, `\\"`, -1)
}

func escapeBreakLine(str string) string {
	return strings.Replace(strings.Replace(str, `\n`, `\\n`, -1), `\t`, `\\t`, -1)
}
