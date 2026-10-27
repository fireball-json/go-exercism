package booking

import "time"

// dateLayouts 是本练习中出现过的所有日期格式。
// Go 用「参考时间」来定义布局：2006=年 01=月 02=日 15=时(24制) 04=分 05=秒
// January=月份全称 Monday=星期全称
var dateLayouts = []string{
	"1/2/2006 15:04:05",                // 7/13/2020 20:32:00
	"January 2, 2006 15:04:05",         // October 3, 2019 20:32:00
	"Monday, January 2, 2006 15:04:05", // Thursday, May 13, 2010 20:32:00
}

// parseDate 尝试用各个已知布局解析日期字符串。
// 因为这几个函数的入参格式并不统一（Schedule/Description 用数字斜杠格式，
// HasPassed 用「月份全称」格式，IsAfternoonAppointment 还带星期），
// 所以统一走这里，逐个布局尝试，解析成功就返回。
func parseDate(date string) time.Time {
	for _, layout := range dateLayouts {
		if t, err := time.Parse(layout, date); err == nil {
			return t
		}
	}
	// 全部布局都失败：返回零值时间（调用方拿到的就是 0001-01-01）
	return time.Time{}
}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	return parseDate(date)
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	// 解析出来的预约时间（UTC）和当前时间比较，预约早于现在即已过去
	return parseDate(date).Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
// 下午的定义：>= 12:00 且 < 18:00
func IsAfternoonAppointment(date string) bool {
	hour := parseDate(date).Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t := parseDate(date)
	// Format 里：Monday=星期全称 January=月份全称 2=日（不带前导 0）2006=年
	// 15:04 是 24 小时制的时:分
	return "You have an appointment on " + t.Format("Monday, January 2, 2006") + ", at " + t.Format("15:04") + "."
}

// AnniversaryDate returns a Time with this year's anniversary.
// 美容院开业日是 2012 年 9 月 15 日，所以每年都是 9 月 15 日。
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
