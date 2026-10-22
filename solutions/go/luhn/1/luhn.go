package luhn

import "strings"

func Valid(id string) bool {
	clearId := strings.ReplaceAll(id, " ", "")
	if len(clearId) <= 1 {
		return false
	}
	nums := []rune(clearId)
	sum := 0
	flag := false
	errorId := false
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > '9' || nums[i] < '0' {
			errorId = true
			break
		}
		digit := int(nums[i] - '0')
		if !flag {
			sum += digit
			flag = true
		} else {
			if digit >= 5 {
				sum += 2*digit - 9
			} else {
				sum += 2 * digit
			}
			flag = false
		}
	}
	if errorId {
		return false
	}
	return sum%10 == 0
}
