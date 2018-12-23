package bob

import (
	"strings"
	"unicode"
)

// Hey triggers a response from Bob
func Hey(remark string) (response string) {

	response = "Whatever."

	if strings.HasSuffix(remark, "?") {
		response = "Sure."
	}
	rune := []rune(remark)
	if strings.HasSuffix(remark, "!") || unicode.IsUpper(rune) {
		response = "Whoa, chill out!"
	}

	return
}
