package main

import "fmt"

func main() {
	var pi = 3
	var a = 10

	var pA = &a
	var pPi = &pi

	fmt.Println(pA, pPi)
	fmt.Println(*pA, *pPi)

	var ppA = &pA
	fmt.Println(ppA, *ppA, **ppA)

	var pppA = &ppA
	fmt.Println(pppA, *pppA, **pppA, ***pppA)

	var p = *&*&*&*&*&a
	fmt.Println(p)

	x, y := 4, 5
	res := 0
	fmt.Println(x, y, res)
	sum(x, y, &res)
	fmt.Println(x, y, res)
}

func sum(a, b int, result *int) {
	*result = a + b
}
