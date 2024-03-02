package main

import "fmt"

func main() {

	//defer func() {
	//	fmt.Println("This is the end of program 1")
	//}()
	//
	//defer func() {
	//	fmt.Println("This is the end of program 2")
	//}()
	//
	//defer func() {
	//	fmt.Println("This is the end of program 3")
	//}()
	//
	//defer func() {
	//	fmt.Println("This is the end of program 4")
	//}()
	//
	//defer func() {
	//	fmt.Println("This is the end of program 5")
	//}()
	//
	//a := 17
	//f := func() {
	//	fmt.Println("hello world")
	//}
	//
	//fmt.Printf("%T, %v\n", a, a)
	//fmt.Printf("%T, %v\n", f, f)
	//
	//for i := 0; i < 1000; i++ {
	//	fmt.Print(i, " ")
	//}

	fmt.Println(sum())
	fmt.Println(sum(1, 2, 3, 1000))

	sum := func(a, b int) int {
		return a + b
	}

	fmt.Println(sum(17, 19))
}

func rectangleCharacteristics(a, b float64) (p float64, s float64) {
	p = 2 * (a + b)
	s = a * b

	return p, s
}

func isAbleToDrive(age int) bool {
	return age > 18
	//if age < 18 {
	//	return false
	//}
	//return true
}

func isAbleToDriveUpdated(age int) (result bool, a int, b float64) {
	if age > 18 {
		result = true
	}

	return
}

// + анонимные функции
// + отложенный запуск (deferred call)
// + функции с неопределенным количеством параметров

func sum(n ...int) (s int) {
	fmt.Println(n)

	for index, element := range n {
		s += element
		fmt.Println("#: ", index, "\tval: ", element)
	}

	return s
}
