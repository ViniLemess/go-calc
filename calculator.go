package main

import (
	"errors"
	"fmt"
)

var history []string

func Sum(x float64, y float64) float64 {
	sum := x + y
	history = append(history, string(fmt.Sprintf("%f+%f=%f", x, y, sum)))
	return sum
}

func Sub(x float64, y float64) float64 {
	difference := x - y
	history = append(history, string(fmt.Sprintf("%f-%f=%f", x, y, difference)))
	return difference
}

func Mul(x float64, y float64) float64 {
	product := x * y
	history = append(history, string(fmt.Sprintf("%fx%f=%f", x, y, product)))
	return product
}

func Div(x float64, y float64) (float64, error) {
	if y == 0 {
		return -1, errors.New("divider cannot be 0")
	}
	quotient := x / y
	history = append(history, string(fmt.Sprintf("%f/%f=%f", x, y, quotient)))
	return quotient, nil
}
