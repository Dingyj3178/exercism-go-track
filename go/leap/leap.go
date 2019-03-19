// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear is calculate whether the year is Leap year
func IsLeapYear(i int) bool {
	result := (i%4 == 0)
	result2 := (i%100 != 0)
	result3 := (i%400 == 0)
	// println(i, result, result2, result3)
	// println(result2)
	// println(result3)

	return (result && result2) || result3
}
