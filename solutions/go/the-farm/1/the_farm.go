package thefarm

import (
	"errors"
	"fmt"
)

// InvalidCowsError is returned when the number of cows is invalid.
type InvalidCowsError struct {
	NumberOfCows int
	Message      string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.NumberOfCows, e.Message)
}

// DivideFood divides the fodder amount by the number of cows.
func DivideFood(f FodderCalculator, n int) (float64, error) {
	a, err1 := f.FatteningFactor()
	b, err2 := f.FodderAmount(n)
	if err1 != nil {
		return 0, err1
	}
	if err2 != nil {
		return 0, err2
	}
	return a * b / float64(n), nil
}

// ValidateInputAndDivideFood divides the fodder only if the number of cows is positive.
func ValidateInputAndDivideFood(f FodderCalculator, n int) (float64, error) {
	if n > 0 {
		return DivideFood(f, n)
	}
	return 0, errors.New("invalid number of cows")
}

// ValidateNumberOfCows returns an error if the number of cows is not positive.
func ValidateNumberOfCows(n int) error {
	switch {
	case n < 0:
		return &InvalidCowsError{NumberOfCows: n, Message: "there are no negative cows"}
	case n == 0:
		return &InvalidCowsError{NumberOfCows: n, Message: "no cows don't need food"}
	default:
		return nil
	}
}
