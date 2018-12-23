package isogram

import (
	"strings"
)

// IsIsogram determines if a word is an isogram or not
func IsIsogram(word string) bool {
	wordCharacters := []rune(strings.ToLower(word))
	uniqueSlice := getUniqueValues(wordCharacters)
	return len(wordCharacters) == len(uniqueSlice)
}

func getUniqueValues(runeSlice []rune) []rune {
	keys := make(map[rune]bool)
	list := []rune{}
	for _, entry := range runeSlice {
		if entry == ' ' || entry == '-' {
			list = append(list, entry)
			continue
		} else if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
