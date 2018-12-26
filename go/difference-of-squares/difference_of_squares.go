package diffsquares

// SquareOfSum returns the the square of the sum of the first N natural numbers.
func SquareOfSum(number int) (square int) {
	for i := 1; i <= number; i++ {
		square += i
	}
	return square * square
}

// SumOfSquares returns the  sum of the squares of the first N natural numbers.
func SumOfSquares(number int) (sum int) {

	for i := 1; i <= number; i++ {
		sum += i * i
	}
	return sum
}

// Difference returns the difference between the square of the sum and the sum of the squares of the first N natural numbers.
func Difference(number int) (difference int) {
	squareOfSum := SquareOfSum(number)
	sumOfSquares := SumOfSquares(number)
	return squareOfSum - sumOfSquares
}
