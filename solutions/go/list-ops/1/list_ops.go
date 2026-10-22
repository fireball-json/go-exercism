package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, cur := range s {
		initial = fn(initial, cur)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := len(s) - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	var res IntList
	for _, cur := range s {
		if fn(cur) {
			res = append(res, cur)
		}
	}
	return res
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	var res IntList
	for _, cur := range s {
		res = append(res, fn(cur))
	}
	return res
}

func (s IntList) Reverse() IntList {
	var res IntList
	for _, cur := range s {
		res = append(IntList{cur}, res...)
	}
	return res
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	res := s
	for _, arr := range lists {
		res = append(res, arr...)
	}
	return res
}
