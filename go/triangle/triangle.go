// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"math"
	"sort"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

func checkzero(a, b, c float64) bool {
	return (a <= 0 || b <= 0 || c <= 0)
}
func checktype(a, b, c float64) bool {
	return (math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) || math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c))
}
func chcektriangle(a, b, c float64) bool {
	boarder := []float64{
		a,
		b,
		c,
	}
	sort.Float64s(boarder)
	return boarder[2] > boarder[0]+boarder[1]
}
func checkeq(a, b, c float64) bool {
	return a == b && a == c
}
func checkIso(a, b, c float64) bool {
	boarder := []float64{
		a,
		b,
		c,
	}
	sort.Float64s(boarder)
	return boarder[0] == boarder[1] || boarder[2] == boarder[1]
}

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	var k Kind
	k = Sca
	switch {
	case checkzero(a, b, c) || chcektriangle(a, b, c) || checktype(a, b, c):
		k = NaT
	case checkeq(a, b, c):
		k = Equ
	case checkIso(a, b, c):
		k = Iso
	}
	return k
}
