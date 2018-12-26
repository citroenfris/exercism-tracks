package proverb

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) (text []string) {
	totalLength := len(rhyme)
	if totalLength <= 0 {
		return text
	}

	for index := range rhyme {
		if index+1 == totalLength {
			text = append(text, "And all for the want of a "+rhyme[0]+".")
			return text
		}
		text = append(text, "For want of a "+rhyme[index]+" the "+rhyme[index+1]+" was lost.")
	}

	return text
}
