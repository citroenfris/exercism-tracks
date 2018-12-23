package gigasecond

import "time"

// AddGigasecond adds 1000000000 seconds
func AddGigasecond(t time.Time) time.Time {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	time := t.Local().Add(1000000000 * time.Second)
	return time
}
