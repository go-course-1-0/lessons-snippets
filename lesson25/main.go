package main

import (
	"errors"
	"fmt"
)

var ErrDivisionByZero = errors.New("division by zero")

func main() {
	fmt.Println(SayHello())
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

func SayHello() string {
	return "Hello"
}
