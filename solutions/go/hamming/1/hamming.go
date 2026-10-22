package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("error string")
	}
	n := 0
	runes_a := []rune(a)
	runes_b := []rune(b)
	for i, cur := range runes_a {
		if cur != runes_b[i] {
			n += 1
		}
	}
	return n, nil
}
