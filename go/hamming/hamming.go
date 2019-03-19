package hamming

import (
	"errors"
)

//Distance is a funtion to calculate hamming distance.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		errMessage := errors.New("not match")
		i := 0
		return i, errMessage
	}
	// fmt.Printf("%v\n", len(a))
	var i int
	for r := 0; r < len(a); r++ {
		// fmt.Printf("%s %s:%v\n", string(a[r]), string(b[r]), r)
		if a[r] != b[r] {
			i++
		}
	}
	return i, nil
}
