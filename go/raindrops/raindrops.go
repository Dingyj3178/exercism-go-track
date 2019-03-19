package raindrops

import "fmt"

//Convert function convert a number to a string
func Convert(a int) string {
	result := ""
	if a%3 == 0 {
		result += "Pling"
	}
	if a%5 == 0 {
		result += "Plang"
	}
	if a%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		return fmt.Sprintf("%d", a)
	}
	return result
}
