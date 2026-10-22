package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	step := 0
	return CollatzConjectureRecursive(n, step)
}

func CollatzConjectureRecursive(n int, step int) (int, error) {
	if n <= 0 {
		return 0, errors.New("wrong number")
	}
	if n == 1 {
		return step, nil
	}
	if n%2 == 0 {
		step += 1
		return CollatzConjectureRecursive(n/2, step)
	} else {
		step += 1
		return CollatzConjectureRecursive(n*3+1, step)
	}
}
