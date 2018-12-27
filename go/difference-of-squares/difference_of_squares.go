package diffsquares

// SquareOfSum returns the the square of the sum of the first N natural numbers.
func SquareOfSum(number int) (square int) {
	sum := (number * (number + 1)) / 2 // note to self instead of /2, >> 1 is also possible.
	return sum * sum
}

// SumOfSquares returns the  sum of the squares of the first N natural numbers.
func SumOfSquares(number int) (sum int) {
	return number * (number + 1) * (2*number + 1) / 6
}

// Difference returns the difference between the square of the sum and the sum of the squares of the first N natural numbers.
func Difference(number int) (difference int) {
	squareOfSum := SquareOfSum(number)
	sumOfSquares := SumOfSquares(number)
	return squareOfSum - sumOfSquares
}
