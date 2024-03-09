package main

import (
	"fmt"
)

func main() {
	//var a [10]int // declare an int array with length 10. Array length is part of the type!
	//a[3] = 42     // set elements
	//i := a[3]     // read elements

	//var students [6]string
	//fmt.Println(students)
	//students[0] = "Siyovush"
	//fmt.Println(students)
	//students[5] = "Farhod"
	//students[3] = "Faridun"
	//fmt.Println(students)
	//
	//fmt.Println(students[5])
	//goodStudent := students[5]
	//fmt.Println(goodStudent)
	//
	//fmt.Println("\nWith a classic loop 1:")
	//for i := 0; i < 6; i++ {
	//	fmt.Println(students[i])
	//	students[i] = "Good guy"
	//}
	//
	//fmt.Println("\nWith a classic loop 2:")
	//for i := 0; i < len(students); i++ {
	//	fmt.Println(students[i])
	//}
	//
	//fmt.Println("\nWith a range loop:")
	//for _, element := range students {
	//	fmt.Println("Some Student:", element)
	//	element = "Bad guy"
	//}
	//fmt.Println(students)
	//
	//// declare and initialize
	//var b = [2]int{1, 2}
	//c := [3]int{1, 2, 3}            //shorthand
	//d := [...]int{1, 2, 3, 4, 5, 6} // elipsis -> Compiler figures out array length
	//
	//var f = [5]int{1, 9}
	//fmt.Println(b, c, d, f)
	//
	//for i := 0; i < 6; i++ {
	//	var name string
	//	fmt.Scan(&name)
	//	students[i] = name
	//}
	//
	//fmt.Println(students)

	//var a []int                               // declare a slice - similar to an array, but length is unspecified
	//var a = []int{1, 2, 3, 4}                 // declare and initialize a slice (backed by the array given implicitly)
	//a := []int{1, 2, 3, 4}                    // shorthand
	//chars := []string{0: "a", 2: "c", 1: "b"} // ["a", "b", "c"]
	//
	//var b = a[lo:hi]     // creates a slice (view of the array) from index lo to hi-1
	//var b = a[1:4]       // slice from index 1 to 3
	//var b = a[:3]        // missing low index implies 0
	//var b = a[3:]        // missing high index implies len(a)
	//a = append(a, 17, 3) // append items to slice a
	//c := append(a, b...) // concatenate slices a and b
	//
	//// create a slice with make
	//a = make([]byte, 5, 5) // first arg length, second capacity
	//a = make([]byte, 5)    // capacity is optional
	//
	//// create a slice from an array
	//x := [3]string{"Лайка", "Белка", "Стрелка"}
	//s := x[:] // a slice referencing the storage of x

	var students []string

	//for {
	//	var command string
	//	fmt.Scan(&command)
	//	if command == "add" {
	//		var name string
	//		fmt.Scan(&name)
	//		students = append(students, name)
	//	} else {
	//		break
	//	}
	//}

	for i := 0; i < 15; i++ {
		students = append(students, fmt.Sprintf("Student %d", i))
	}

	fmt.Println(students)

	fmt.Println()

	fmt.Println(students[5:10])

	fmt.Println()

	PrintBeautiful(&students)

	fmt.Println()

	fmt.Println(students)

	fmt.Println()

	PrintBeautiful2(students)

	fmt.Println()

	fmt.Println(students)

	fmt.Println()

	ar := [3]string{"First", "Second", "Third"}

	PrintBeautiful3(ar)

	fmt.Println()

	fmt.Println(ar)

	fmt.Println()
}

func PrintBeautiful(sl *[]string) {
	fmt.Println("PrintBeautiful ===")
	fmt.Printf("%T\n", sl)
	fmt.Println(sl)
	*sl = append(*sl, "New Student")
	fmt.Println("===")
}

func PrintBeautiful2(sl []string) {
	fmt.Println("PrintBeautiful2 ===")
	fmt.Printf("%T\n", sl)
	fmt.Println(sl)
	sl[0] = "I changed it!"
	fmt.Println("===")
}

func PrintBeautiful3(ar [3]string) {
	fmt.Println("PrintBeautiful3 ===")
	fmt.Printf("%T\n", ar)
	fmt.Println("Before change:")
	fmt.Println(ar)

	ar[0] = "I changed it!"

	fmt.Println("After change:")
	fmt.Println(ar)

	fmt.Println("===")
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}
