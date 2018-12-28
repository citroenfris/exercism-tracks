package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(ccNumber string) (valid bool) {
	if len(ccNumber) <= 1 {
		return false
	}

	ccNumber = trimAllWhitespaces(ccNumber)
	ccInt, err := strconv.Atoi(ccNumber)
	if err != nil {
		return false
	}

	if ccInt <= 0 && len(ccNumber) < 2 {
		return false
	}

	cardDigits := []int{}
	for ccInt > 0 {
		cardDigits = append(cardDigits, ccInt%10)
		ccInt = ccInt / 10
	}

	for i := 1; i < len(cardDigits); i++ {
		if i%2 == 0 {
			continue
		}

		newValue := cardDigits[i] * 2

		if newValue > 9 {
			newValue -= 9
		}
		cardDigits[i] = newValue
	}

	var sum int
	for _, digit := range cardDigits {
		sum += digit
	}

	return sum%10 == 0
}

func trimAllWhitespaces(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
