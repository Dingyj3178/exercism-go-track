package raindrops

import (
	"strconv"
)

//Convert function convert a number to a string
func Convert(a int) string {
	factor := make([]string, 10)
	result := ""
	factor[3] = "Pling"
	factor[5] = "Plang"
	factor[7] = "Plong"
	for i, f := range factor {
		if i > 0 && a%i == 0 {
			result += f
		}
	}
	if result == "" {
		result = strconv.Itoa(a)
	}
	return result
}
