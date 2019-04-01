package accumulate

// Accumulate change slice content base on the converter
func Accumulate(s []string, converter func(string) string) []string {
	result := s
	for i, r := range s {
		result[i] = converter(r)
	}
	return result
}
