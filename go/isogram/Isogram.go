package isogram

import "strings"

func noSpaceorHyphens(s string) string {
	rep := strings.NewReplacer(" ", "", "-", "")
	return rep.Replace(s)
}

// IsIsogram is a function to check string duplicate
func IsIsogram(input string) bool {
	input = noSpaceorHyphens(input)
	input = strings.ToLower(input)
	temp := map[string]struct{}{}
	for _, r := range input {
		if _, ok := temp[string(r)]; !ok {
			temp[string(r)] = struct{}{}
		} else {
			return false
		}
	}
	return true
}
