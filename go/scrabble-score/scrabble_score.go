package scrabble

import (
	"strings"
)

var scoreMapping = map[rune]int{
	65: 1,
	69: 1,
	73: 1,
	79: 1,
	85: 1,
	76: 1,
	78: 1,
	82: 1,
	83: 1,
	84: 1,
	68: 2,
	71: 2,
	66: 3,
	67: 3,
	77: 3,
	80: 3,
	70: 4,
	72: 4,
	86: 4,
	87: 4,
	89: 4,
	75: 5,
	74: 8,
	88: 8,
	81: 10,
	90: 10,
}

// Score function calculate string parameter's point base on scoreMapping
func Score(input string) int {
	result := 0
	inputUpper := strings.ToUpper(input)
	for _, s := range inputUpper {
		result += scoreMapping[s]
	}
	return result
}
