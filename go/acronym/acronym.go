package acronym

import (
	"strings"
)

// Abbreviate e.g. Ruby on Rails as ROR
func Abbreviate(s string) (out string) {

	s = strings.Replace(s, "-", " ", -1)

	words := strings.Fields(s)

	for i := range words {
		out += string(words[i][0])
	}
	return strings.ToUpper(out)
}
