package etl

import "strings"

func Transform(input map[int][]string) map[string]int {
	result := make(map[string]int)
	for i1, s1 := range input {
		for _, s2 := range s1 {
			result[strings.ToLower(s2)] = i1
		}
	}
	return result
}
