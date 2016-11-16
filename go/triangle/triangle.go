package triangle

import (
	"math"
	"sort"
)

const testVersion = 3

//Kind describes the type of triangle
type Kind int

const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

//KindFromSides returns the Kind of a triangle given the lengths of its sides
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	sides := []float64{a, b, c}
	sort.Float64s(sides)
	x, y, z := sides[0], sides[1], sides[2]

	if z > (x+y) ||
		allAreZero(x, y, z) ||
		anyAreNaN(x, y, z) ||
		anyAreInf(x, y, z) {
		return NaT
	}

	if x == y && y == z {
		return Equ
	}

	if x != y && y != z && z != x {
		return Sca
	}

	if (x == y && y != z) ||
		(x != y && y == z) ||
		(x != y && x == z) {
		return Iso
	}

	return k
}

func anyAreNaN(nums ...float64) bool {
	for _, i := range nums {
		if math.IsNaN(i) {
			return true
		}
	}
	return false
}

func allAreZero(nums ...float64) bool {
	for _, i := range nums {
		if i != 0 {
			return false
		}
	}
	return true
}

func anyAreInf(nums ...float64) bool {
	for _, i := range nums {
		if math.IsInf(i, 0) {
			return true
		}
	}
	return false
}