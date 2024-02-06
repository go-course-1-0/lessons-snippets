package main

import "fmt"

func main() {
	var (
		x1, y1 int
		x2, y2 int
		x3, y3 int
	)

	fmt.Scan(&x1, &y1)
	fmt.Scan(&x2, &y2)
	fmt.Scan(&x3, &y3)

	fmt.Println(x1, y3)
}
