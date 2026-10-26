package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	p := SplitSimple(phrase)
	m := make(map[string]int)
	for _, r := range p {
		m[string(r)] += 1
	}
	return m
}

func SplitSimple(s string) []string {
	var runes []rune
	rs := []rune(s)
	for idx, r := range rs {
		// 只保留字母、数字、空格
		if isok(r) {
			runes = append(runes, r)
		} else {
			prevIsNotSpace := (idx > 0) && rs[idx-1] != ' '
			nextIsNotSpace := (idx < len(rs)-1) && rs[idx+1] != ' '
			if prevIsNotSpace && nextIsNotSpace && r == '\'' {
				runes = append(runes, r)
			} else {
				runes = append(runes, ' ')
			}
		}
	}
	clean := strings.ToLower(string(runes))
	return strings.Fields(clean)
}

func isok(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
