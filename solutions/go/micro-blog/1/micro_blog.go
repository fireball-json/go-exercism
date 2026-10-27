package microblog

// Truncate 把输入截断为最多 5 个字符。
func Truncate(phrase string) string {
	const limit = 5
	// 转成 rune 切片后，len() 数的是字符数而不是字节数
	runes := []rune(phrase)
	if len(runes) <= limit {
		return phrase
	}
	return string(runes[:limit])
}
