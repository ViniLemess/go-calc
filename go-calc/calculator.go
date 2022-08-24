package main

import (
	"errors"
	"fmt"
)

var history []string

func addHistory(x float64, y float64, res float64, op string) {
	history = append(history, fmt.Sprintf("%s%s%s=%s", fmt.Sprint(x), op, fmt.Sprint(y), fmt.Sprint(res)))
}

func calculate(op string, x float64, y float64) (float64, error) {
	var result float64
	switch op {
	case "sum":
		sum := x + y
		addHistory(x, y, sum, "+")
		result = sum
	case "sub":
		difference := x - y
		addHistory(x, y, difference, "-")
		result = difference
	case "mul":
		product := x * y
		addHistory(x, y, product, "*")
		result = product
	case "div":
		if y == 0 {
			return 0, errors.New("illegal division operation, divider cannot be 0")
		}
		quotient := x / y
		addHistory(x, y, quotient, "/")
		result = quotient
	default:
		return 0, errors.New("unsupported operation : " + op)
	}
	return result, nil
}
