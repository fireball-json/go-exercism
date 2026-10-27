package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
// 从左到右扫描日志，遇到的第一个应用标识字符就决定来源；
// 一个都没有则返回 "default"。
func Application(log string) string {
	for _, r := range log {
		switch r {
		case '❗': // U+2757
			return "recommendation"
		case '🔍': // U+1F50D
			return "search"
		case '☀': // U+2600
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	// string(rune) 得到该字符的 UTF-8 编码，ReplaceAll 按字符串替换全部出现处
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	// 必须按 rune 计数而不是用 len()：
	// len() 返回字节数，而 rune 最多占 4 字节（emoji 就是 4 字节），
	// 例如 "exercism❗" 有 9 个字符但占 11 个字节。
	return utf8.RuneCountInString(log) <= limit
}
