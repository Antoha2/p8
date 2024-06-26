package main

import "errors"

func z9(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("на ноль делить нельзя")
	}
	return a / b, nil
}
