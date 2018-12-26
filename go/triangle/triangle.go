package triangle

import (
	"math"
)

// Kind is an enumerated type for triangle kinds
type Kind int

// Types of triangles (an enumareted list)
const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

/*KindFromSides determines the type of triangle formed by sides a, b and c
the triangle is invaid if the longest side is equal or longer than the shorter sides
or any of the sides are invalid (0 or less, NaN or Inf+)*/
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	aEqualsB := a == b
	bEqualsC := b == c
	aEqualsC := a == c
	if !isTriangle(a, b, c) {
		k = NaT
	} else if aEqualsB && bEqualsC {
		k = Equ
	} else if !aEqualsB && !bEqualsC && !aEqualsC {
		k = Sca
	} else if aEqualsB || bEqualsC || aEqualsC {
		k = Iso
	}

	return k
}

func isTriangle(a, b, c float64) bool {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) || a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	if a+b < c || c+b < a || a+c < b {
		return false
	}

	return true
}
