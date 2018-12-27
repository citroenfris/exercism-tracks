package reverse

import (
	"unicode/utf8"
)

// String returns the input as reversed string
func String(str string) (result string) {
	for len(str) > 0 {
		r, size := utf8.DecodeLastRuneInString(str)
		str = str[:len(str)-size]
		result = result + string(r)
	}
	return result
}
