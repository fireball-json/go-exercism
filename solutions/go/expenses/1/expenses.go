package expenses

import (
	"errors"
	"slices"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	res := []Record{}
	for _, v := range in {
		if predicate(v) {
			res = append(res, v)
		}
	}
	return res
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		if r.Day <= p.To && r.Day >= p.From {
			return true
		} else {
			return false
		}
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		if r.Category == c {
			return true
		}
		return false
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	arr := Filter(in, ByDaysPeriod(p))
	count := 0.0
	for _, v := range arr {
		count += v.Amount
	}
	return count
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	isC := slices.ContainsFunc(in, func(v Record) bool {
		if v.Category == c {
			return true
		}
		return false
	})
	if !isC {
		return 0, errors.New("unknown category entertainment")
	}
	arr := Filter(in, ByDaysPeriod(p))
	arr = Filter(arr, ByCategory((c)))
	count := 0.0
	for _, v := range arr {
		count += v.Amount
	}
	return count, nil
}
