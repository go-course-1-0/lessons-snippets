package main

import "fmt"

func main() {
	hour := 0
	fmt.Println("Который сейчас час?")
	fmt.Scan(&hour)

	//if hour == 18 {
	//	fmt.Println("Ваш урок начался")
	//} else if hour >= 20 {
	//	fmt.Println("Ваш урок закончился")
	//} else if hour > 18 {
	//	fmt.Println("Ваш урок уже идет")
	//} else {
	//	fmt.Println("Ваш урок еще не начался")
	//}

	// 18 - 20 дарс рафта истодааст
	// < 18 - дарс сар нашудааст, то дарса 18-hour
	// < 18 -дарс сар нашудааст, лекин 1 соат пас сар мешавад, бодиккат бошед
	// > 20 - дарс тамом шудааст

	if hour < 0 || hour > 23 {
		fmt.Println("Введите число от 0 до 23")
		return
	}

	if hour >= 18 && hour < 20 {
		fmt.Println("дарс рафта истодааст")
	} else if hour < 18 {
		left := 18 - hour
		fmt.Println("дарс сар нашудааст, то дарса", left, "соат мондааст")
		if left <= 1 {
			fmt.Println("бодиккат бошед!")
		}
	} else {
		fmt.Println("дарс тамом шудааст")
	}

	f()

	fmt.Println("Программа завершилась")

}

func f() {
	//
	//
	//
	return
	//
	//
	//
}
