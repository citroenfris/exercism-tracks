package main

import (
	"fmt"
	"strconv"
)

func main() {
	var _ = Valid("")
}
func Valid(ccNumber string) (valid bool) {
	// ccNumber = "378282246310005"
	ccNumber = "4539148803436467"
	ccInt, err := strconv.Atoi(ccNumber)
	if err != nil {
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

	fmt.Printf("Result: %v \n", cardDigits)
	fmt.Printf("Sum: %v \n", sum)

	return sum%10 == 0
}
