package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//var a, b int
	//fmt.Scan(&a, &b)
	//if b < a {
	//	// a -> b
	//	// b -> a
	//	//temp := a
	//	//a = b
	//	//b = temp
	//	a, b = b, a
	//}

	for {
		var command bool
		fmt.Print("Continue? > ")
		fmt.Scan(&command)
		if !command {
			break
		}

		l := rand.Intn(20)
		fmt.Println()
		fmt.Printf("New array with length %d\n", l)
		var numbers []int
		for i := 0; i < l; i++ {
			number := rand.Intn(100)
			numbers = append(numbers, number)
		}

		fmt.Println(numbers)
		fmt.Println()

		fmt.Println()
		fmt.Println("Searching for zero:")

		var found bool
		var indexOfZero int
		for i := 0; i < len(numbers); i++ {
			fmt.Printf("i: %d; ", i)

			if numbers[i] == 0 {
				found = true
				indexOfZero = i
				break
			}
		}
		fmt.Println()

		if found {
			fmt.Println("Hooray! We found zero! It is under the index of:", indexOfZero)
		} else {
			fmt.Println("Unfortunately, there is no zero in your array :(")
		}

		fmt.Println()
		fmt.Println("Sorting by ascending:")
		// Bubble sort
		for j := 0; j < len(numbers); j++ {
			for i := 0; i < len(numbers)-1-j; i++ {
				fmt.Printf("i: %d; ", i)
				if numbers[i] > numbers[i+1] {
					numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				}
			}
			fmt.Println()
		}

		fmt.Println("I sorted it by ascending values:")
		fmt.Println(numbers)

		fmt.Println()
		fmt.Println("Sorting by descending:")
		// Bubble sort
		for j := 0; j < len(numbers); j++ {
			for i := 0; i < len(numbers)-1-j; i++ {
				fmt.Printf("i: %d; ", i)
				if numbers[i] < numbers[i+1] {
					numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				}
			}
			fmt.Println()
		}

		fmt.Println("I sorted it by descending values:")
		fmt.Println(numbers)
	}
}
