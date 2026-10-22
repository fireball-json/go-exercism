package sumofmultiples

import (
	"math"
	"sort"
)

func SumMultiples(limit int, divisors ...int) int {
	arr := []int{}
	for _, val := range divisors {
		if val == 0 {
			break
		}
		j := math.Floor(float64((limit - 1) / val))
		for i := 1; i <= int(j); i++ {
			arr = append(arr, i*val)
		}
	}
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	j := 0
	for i := 1; i < len(arr); i++ {
		if arr[j] != arr[i] {
			j++
			arr[j] = arr[i]
		}
	}
	arr = arr[:j+1]
	res := 0
	for _, v := range arr {
		res += v
	}
	return res
}
