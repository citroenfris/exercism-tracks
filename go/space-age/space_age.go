package space

// Age returns how old someone would be on the given planet
func Age(seconds float64, planet string) float64 {
	var ratio float64
	earth := 31557600.00

	switch planet {
	case "Earth":
		ratio = 1
	case "Mercury":
		ratio = 0.2408467
	case "Venus":
		ratio = 0.61519726
	case "Mars":
		ratio = 1.8808158
	case "Jupiter":
		ratio = 11.862615
	case "Saturn":
		ratio = 29.447498
	case "Uranus":
		ratio = 84.016846
	case "Neptune":
		ratio = 164.79132
	}

	return seconds / (ratio * earth)
}
