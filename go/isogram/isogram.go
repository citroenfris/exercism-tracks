package isogram

import (
	"strings"
)

// IsIsogram determines if a word is an isogram or not
// It compares the length of a string together with the length of the unique values of that string.
// If the length of the word is longer than the count of unique values it is not an isogram
func IsIsogram(word string) bool {
	wordCharacters := []rune(strings.ToLower(word))
	keys := make(map[rune]bool)
	for _, char := range wordCharacters {
		if char == ' ' || char == '-' {
			continue
		}
		if keys[char] {
			return false
		}
		keys[char] = true
	}
	return true
}
