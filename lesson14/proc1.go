package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64

	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		PowerA3(a, &b)
		fmt.Println(b)
	}
}

func PowerA3(a float64, b *float64) {
	*b = math.Pow(a, 3)
}
