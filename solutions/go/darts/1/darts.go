package darts

import "math"

func Score(x, y float64) int {
	r := math.Sqrt(x*x + y*y)
	if r <= 1 {
		return 10
	} else if r <= 5 {
		return 5
	} else if r <= 10 {
		return 1
	} else {
		return 0
	}
}
