package meetup

import "time"

// WeekSchedule 表示一个月中的第几个星期几。
// 用 int 底层的自定义类型，这样可以做 switch 判断，也能作为 map 的 key。
type WeekSchedule int

const (
	// First 表示当月第一个该星期几（1~7 号）
	First WeekSchedule = iota
	// Second 表示当月第二个该星期几（8~14 号）
	Second
	// Third 表示当月第三个该星期几（15~21 号）
	Third
	// Fourth 表示当月第四个该星期几（22~28 号）
	Fourth
	// Teenth 表示当月 teenth 日（13~19 号），每个月每个星期几恰好有一个
	Teenth
	// Last 表示当月最后一个该星期几
	Last
)

// Day 返回指定月份、年份中，符合给定周次和星期几的日期（几号）。
func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	switch wSched {
	case First:
		return firstMatchDay(wDay, month, year) + 0
	case Second:
		return firstMatchDay(wDay, month, year) + 7
	case Third:
		return firstMatchDay(wDay, month, year) + 14
	case Fourth:
		return firstMatchDay(wDay, month, year) + 21
	case Teenth:
		return teenthDay(wDay, month, year)
	case Last:
		return lastMatchDay(wDay, month, year)
	default:
		return 0
	}
}

// firstMatchDay 返回当月第一次出现的 wDay 是几号（结果落在 1~7）。
func firstMatchDay(wDay time.Weekday, month time.Month, year int) int {
	first := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	// 首日星期与目标星期相差几天（取模保证为 0~6 的正数）
	offset := (int(wDay) - int(first.Weekday()) + 7) % 7
	return 1 + offset
}

// teenthDay 返回当月 teenth 日（13~19 号）里星期为 wDay 的那一天。
// 做法：从 13 号起算，13~19 号这七天恰好覆盖一周，
// 所以其中必定有且仅有一天是 wDay。
func teenthDay(wDay time.Weekday, month time.Month, year int) int {
	thirteenth := time.Date(year, month, 13, 0, 0, 0, 0, time.UTC)
	offset := (int(wDay) - int(thirteenth.Weekday()) + 7) % 7
	return 13 + offset
}

// lastMatchDay 返回当月最后一个 wDay 是几号。
// 做法：先取当月最后一天，再往前回退到最近的那个 wDay。
func lastMatchDay(wDay time.Weekday, month time.Month, year int) int {
	// month+1 的 0 号 = 本月的最后一天
	last := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
	back := (int(last.Weekday()) - int(wDay) + 7) % 7
	return last.Day() - back
}
