package isbnverifier

func IsValidISBN(isbn string) bool {
	index := 10
	sum := 0
	errorIsbn := false
	for idx, cur := range isbn {
		if idx == len(isbn)-1 {
			if cur <= '9' && cur >= '0' {
				sum += int(cur - '0')
			} else if cur == 'X' {
				sum += 10
			} else {
				errorIsbn = true
				break
			}
		} else {
			if cur <= '9' && cur >= '0' {
				sum += int(cur-'0') * index
				index -= 1
			} else if cur != '-' {
				errorIsbn = true
				break
			}
		}
	}
	if errorIsbn || index != 1 {
		return false
	}
	return sum%11 == 0
}
