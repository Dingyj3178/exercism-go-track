package raindrops

import "strconv"

//Convert function convert a number to a string.
func Convert(a int) string {
	switch {
	case a%105 == 0:
		return "PlingPlangPlong"
	case a%35 == 0:
		return "PlangPlong"
	case a%21 == 0:
		return "PlingPlong"
	case a%15 == 0:
		return "PlingPlang"
	case a%3 == 0:
		return "Pling"
	case a%5 == 0:
		return "Plang"
	case a%7 == 0:
		return "Plong"
	default:
		return strconv.Itoa(a)
	}
}
