package myMath

import "errors"

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func Subtract(a, b int) int {
	return b - a
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by Zero")
	}
	return a / b, nil
}
