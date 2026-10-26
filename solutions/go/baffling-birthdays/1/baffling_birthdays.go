package bafflingbirthdays

import (
	"fmt"
	"math/rand"
	"time"
)

func SharedBirthday(dates []time.Time) bool {
	hashMap := make(map[any]int)
	for _, v := range dates {
		month := v.Month()
		day := v.Day()
		key := fmt.Sprintf("%02d", int(month)) + fmt.Sprintf("%02d", int(day))
		hashMap[key]++
		if hashMap[key] >= 2 {
			return true
		}
	}
	return false
}

func RandomBirthdates(size int) []time.Time {
	arr := make([]time.Time, size)
	for i, _ := range arr {
		y := rand.Intn(9999)
		for ifLeapYear(y) {
			y = rand.Intn(9999)
		}
		m := time.Month(rand.Intn(12) + 1)
		maxD := time.Date(y, time.Month(m+1), 0, 0, 0, 0, 0, time.Local).Day()
		d := rand.Intn(maxD) + 1
		arr[i] = time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	}
	return arr
}

func EstimatedProbability(size int) float64 {
	p1 := float64(1)
	for i := 0; i < size; i++ {
		p1 = p1 * float64((365 - i)) / 365
	}
	return (1 - p1) * 100
}

func ifLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}
