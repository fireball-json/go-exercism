package parsinglogfiles

import (
	"regexp"
)

var (
	validLineRe = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	separatorRe = regexp.MustCompile(`<[~*=-]*>`)
	quotedPwRe  = regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
	eolRe       = regexp.MustCompile(`end-of-line\d+`)
	userRe      = regexp.MustCompile(`User\s+(\S+)`)
)

// IsValidLine 判断日志行是否以合法的级别标记开头。
func IsValidLine(text string) bool {
	return validLineRe.MatchString(text)
}

// SplitLogLine 按 <--> / <=> 这类分隔符切分日志行。
func SplitLogLine(text string) []string {
	return separatorRe.Split(text, -1)
}

// CountQuotedPasswords 统计引号中出现 password 的行数。
func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		if quotedPwRe.MatchString(line) {
			count++
		}
	}
	return count
}

// RemoveEndOfLineText 移除行内残留的 end-of-line 加行号文本。
func RemoveEndOfLineText(text string) string {
	return eolRe.ReplaceAllString(text, "")
}

// TagWithUserName 给包含 "User <name>" 的行加上 [USR] <name> 标签。
func TagWithUserName(lines []string) []string {
	tagged := make([]string, 0, len(lines))
	for _, line := range lines {
		match := userRe.FindStringSubmatch(line)
		if match == nil {
			tagged = append(tagged, line)
			continue
		}
		tagged = append(tagged, "[USR] "+match[1]+" "+line)
	}
	return tagged
}
