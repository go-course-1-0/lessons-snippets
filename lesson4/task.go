package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	//fmt.Println((a%2 == 0 && b%2 == 0) || (a%2 != 0 && b%2 != 0))
	//fmt.Println((a+b)%2 == 0)
	fmt.Println(a%2 == b%2)
}
