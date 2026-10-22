package clock

import (
	"fmt"
	"math"
)

// Define the Clock type here.
type Clock struct {
	total int
}

func New(h, m int) Clock {
	curM := m + 60*h
	if curM < 0 {
		return Clock{total: curM%1440 + 1440}
	} else {
		return Clock{total: curM % 1440}
	}

}

func (c Clock) Add(m int) Clock {
	curM := c.total + m
	if curM < 0 {
		c.total = 1440 + curM%1440
	} else {
		c.total = curM % 1440
	}
	return c
}

func (c Clock) Subtract(m int) Clock {
	curM := c.total - m
	if curM < 0 {
		c.total = 1440 + curM%1440
	} else {
		c.total = curM % 1440
	}
	return c
}

func (c Clock) String() string {
	trueM := c.total % 60
	trueH := int(math.Floor(float64(c.total / 60)))
	return fmt.Sprintf("%02d:%02d", trueH, trueM)
}
