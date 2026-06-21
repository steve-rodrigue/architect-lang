package antlr

import (
	"strings"
)

func unquote(value string) string {
	return strings.Trim(value, `"`)
}
