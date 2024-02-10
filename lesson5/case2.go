package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)

	switch k {
	case 1:
		fmt.Println("Плохо")
	case 2:
		fmt.Println("Неудовлетворительно")
	case 3:
		fmt.Println("Удовлетворительно")
	case 4:
		fmt.Println("Хорошо")
	case 5:
		fmt.Println("Отлично")
	default:
		fmt.Println("Либо вы очень хорошо учили, либо вы ввели оценку неправильно")
	}
}
