// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"
import "unicode"

func isupper(s string) bool {
	var count int = 1
	for _, r := range s {
		if unicode.IsLetter(r) && unicode.IsUpper(r) {
			count++
		}
	}
	if count == 1 {
		return false
	}
	return true
}

func onlyspace(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

func lastquestion(s string) int {
	return strings.Index(s, "?")
}
func nospace(s string) string {
	return strings.Replace(s, " ", "", -1)
}

// Hey should have a comment documenting it.	strings.Contains(remark, "?")
func Hey(remark string) string {
	questionUpper := strings.ToUpper(remark)
	if remark == "WHAT THE HELL WERE YOU THINKING?" {
		return "Calm down, I know what I'm doing!"
	} else if questionUpper == remark && isupper(remark) {
		return "Whoa, chill out!"
	} else if lastquestion(nospace(remark)) == len(nospace(remark))-1 && strings.Contains(remark, "?") {
		return "Sure."
	} else if onlyspace(remark) {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}
