package main

import "fmt"

// Get a number
// Check if number^2 делится на 3 без остатка

func main() {
	var n int
	fmt.Scan(&n)

	if n2 := n * n; n2%3 == 0 {
		fmt.Println(n2, true)
		return
	}

	fmt.Println(false)

	someFunction()
}

// Variables lifecycle

func someFunction() bool {
	return true
}
