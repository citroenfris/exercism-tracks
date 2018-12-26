package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate will create an acronym of the given string
// hyphenated words are treated as separate words
func Abbreviate(s string) (acronym string) {
	var abbreviation []byte
	s = strings.Replace(s, "-", " ", -1)
	s = strings.ToUpper(s)
	trimmedText := strings.Split(s, " ")

	for _, char := range trimmedText {
		if len(char) <= 0 || !unicode.IsLetter(rune(char[0])) {
			continue
		}
		abbreviation = append(abbreviation, char[0])

	}

	return string(abbreviation)
}
