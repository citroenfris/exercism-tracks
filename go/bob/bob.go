package bob

import (
	"strings"
)

// Hey triggers a response from Bob
func Hey(s string) (response string) {
	s = strings.TrimSpace(s)
	isQuestion := strings.HasSuffix(s, "?")
	if len(s) == 0 {
		return "Fine. Be that way!"
	} else if isYelling(s) && isQuestion {
		return "Calm down, I know what I'm doing!"
	} else if isYelling(s) {
		return "Whoa, chill out!"
	} else if isQuestion {
		return "Sure."
	}
	return "Whatever."
}

func isYelling(s string) bool {
	return s == strings.ToUpper(s) && s != strings.ToLower(s)
}
