package armstrongnumbers

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	s := strconv.Itoa(n)
	l := len(s)
	res := 0
	for _, r := range s {
		d := int(r - '0')
		res += int(math.Pow(float64(d), float64(l)))
	}
	return n == res
}
