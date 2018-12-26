package isogram

import (
	"strings"
)

// IsIsogram determines if a word is an isogram or not
// It compares the length of a string together with the length of the unique values of that string.
// If the length of the word is longer than the count of unique values it is not an isogram
func IsIsogram(word string) bool {
	wordCharacters := []rune(strings.ToLower(word))
	uniqueLetters := getUniqueLetters(wordCharacters)
	return len(wordCharacters) == len(uniqueLetters)
}

// getUniqueValues returns all unique characters in the given slice
// Whitespace value and - are not required to be unique.
func getUniqueLetters(inputSlice []rune) []rune {
	keys := make(map[rune]bool)
	uniqueLetters := []rune{}
	for _, char := range inputSlice {
		// These characters are allowed to duplicate so add them to the uniqueLetters slice
		if char == ' ' || char == '-' {
			uniqueLetters = append(uniqueLetters, char)
			continue
		}
		// If the char is not yet present in the uniqueLetters slice map it can be added to the uniqueLetters slice.
		if _, value := keys[char]; !value {
			keys[char] = true
			uniqueLetters = append(uniqueLetters, char)
		}
	}
	return uniqueLetters
}
