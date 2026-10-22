package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	res := map[string]int{}
	for k, v := range in {
		for _, cur := range v {
			res[strings.ToLower(cur)] = k
		}
	}
	return res
}
