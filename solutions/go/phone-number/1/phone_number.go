package phonenumber

import (
	"errors"
	"fmt"
	"strings"
)

func Number(phoneNumber string) (string, error) {
	res := ""
	for _, cur := range phoneNumber {
		if cur >= '0' && cur <= '9' {
			res += string(cur)
		}
	}
	if len(res) > 10 {
		res = strings.TrimPrefix(res, "1")
	}
	if len(res) != 10 {
		return res, errors.New(fmt.Sprint("err: ", res))
	} else {
		errorNum := false
		for idx, cur := range res {
			if (idx == 0 || idx == 3) && !(cur >= '2' && cur <= '9') {
				errorNum = true
				break
			}
		}
		if errorNum {
			return res, errors.New("error")
		} else {
			return res, nil
		}
	}
}

func AreaCode(phoneNumber string) (string, error) {
	res, err := Number(phoneNumber)
	if err != nil {
		return res, err
	}
	area := ""
	for inx, cur := range res {
		if inx < 3 {
			area += string(cur)
		} else {
			break
		}
	}
	return area, nil
}

func Format(phoneNumber string) (string, error) {
	res, err := Number(phoneNumber)
	if err != nil {
		return res, err
	}
	area := ""
	number := ""
	for inx, cur := range res {
		if inx < 3 {
			area += string(cur)
		} else {
			number += string(cur)
			if inx == 5 {
				number += "-"
			}
		}
	}
	return fmt.Sprintf("(%s) %s", area, number), nil
}
