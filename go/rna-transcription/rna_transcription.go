package strand

func ToRNA(dna string) (rna string) {

	var keyMap = map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	for _, char := range dna {
		rna = rna + string(keyMap[char])
	}
	return rna
}
