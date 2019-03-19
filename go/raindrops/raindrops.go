package raindrops

import (
	"sort"
	"strconv"
)

//Convert function convert a number to a string
func Convert(a int) string {
	result := ""
	keys := []int{3, 5, 7}
	factor := map[int]string{3: "Pling", 5: "Plang", 7: "Plong"}
	sort.Ints(keys)
	for _, f := range keys {
		if a%f == 0 {
			result += factor[f]
		}
	}
	if result == "" {
		result = strconv.Itoa(a)
	}
	return result
}
