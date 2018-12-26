package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	// 	on every year that is evenly divisible by 4
	//   except every year that is evenly divisible by 100
	//     unless the year is also evenly divisible by 400

	isDividableBy4 := year%4 == 0
	isDividableBy100 := year%100 == 0
	isDividableBy400 := year%400 == 0

	if !isDividableBy4 {
		return false
	}
	if isDividableBy100 && !isDividableBy400 {
		return false
	}
	return true
}
