package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var res []string
	for _, cur := range candidates {
		if !strings.EqualFold(subject, cur) {
			isSame := Compare(cur, subject)
			if isSame {
				res = append(res, cur)
			}
		}
	}
	return res
}

func Compare(strA string, strB string) bool {
	if len(strA) != len(strB) {
		return false
	}
	rsA := []rune(strings.ToLower(strA))
	rsB := []rune(strings.ToLower(strB))
	sort.Slice(rsA, func(i, j int) bool { return rsA[i] < rsA[j] })
	sort.Slice(rsB, func(i, j int) bool { return rsB[i] < rsB[j] })
	res := true
	for i := 0; i < len(rsA); i++ {
		if rsA[i] != rsB[i] {
			res = false
			break
		}
	}
	return res
}
