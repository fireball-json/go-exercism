package scrabblescore

import (
	"slices"
	"strings"
)

func Score(word string) int {
	count := 0
	hashMap1 := []string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}
	hashMap2 := []string{"D", "G"}
	hashMap3 := []string{"B", "C", "M", "P"}
	hashMap4 := []string{"F", "H", "V", "W", "Y"}
	hashMap5 := []string{"K"}
	hashMap6 := []string{"J", "X"}
	hashMap7 := []string{"Q", "Z"}
	for _, r := range word {
		w := strings.ToUpper(string(r))
		if slices.Contains(hashMap1, w) {
			count += 1
		} else if slices.Contains(hashMap2, w) {
			count += 2
		} else if slices.Contains(hashMap3, w) {
			count += 3
		} else if slices.Contains(hashMap4, w) {
			count += 4
		} else if slices.Contains(hashMap5, w) {
			count += 5
		} else if slices.Contains(hashMap6, w) {
			count += 8
		} else if slices.Contains(hashMap7, w) {
			count += 10
		}
	}
	return count

}
