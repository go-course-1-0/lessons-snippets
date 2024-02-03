package main

import "fmt"

func main() {
	// general form: var name type = expression

	//var username string = "Siyovush"
	//var username = "Siyovush"
	//username := "Siyovush"

	//var username, password string
	//	var username, password, phone, email = `Siyovush
	//		srgsrg`, "password", 987654321, `email
	//
	//	rsignsrng
	//grgsgs
	//srgjirsjg
	//rsojgpisr
	//`

	//username, password, phone, email := "1", 2.46, 3, "4"

	//var (
	//	username, password string
	//	phone              int = 66 + 44
	//	email                  = "admin@mail.com"
	//)
	//
	////var isRegistered bool
	////postsCount := 0
	//
	//fmt.Println(username, password, phone, email)

	//a := 0
	//a++ // increment
	//a-- // decrement
	//a = a + 50
	//a += 50
	//a -= 50

	// general form for constants
	// const name type = expression
	const pi = 3.14
	const RightAngleDegree int = 90
	const NumberOfMonths = 12

	//const Jan = 1
	//const Feb = 2
	//const Mar = 3

	//const (
	//	Jan = 1
	//	Feb = 2
	//	Mar = 3
	//)

	//const (
	//	Jan = iota % 2 * 5
	//	Feb
	//	Mar
	//	Apr
	//	May
	//	Jun
	//	Jul
	//	Aug
	//	Sep
	//	Oct
	//	Nov
	//	Dec
	//)
	//RightAngleDegree = 40

	//bytesCount, _ := fmt.Println(Jan, Feb, Mar)

	//var variable int = 1000
	//_ = variable
	//
	//fmt.Println(false)

	var result bool

	//age := 0
	//fmt.Println("Введите ваш возраст:")
	//fmt.Scan(&age)

	//result = age >= 18
	//result = age > 17

	//result = age >= 18 && age <= 90
	brand1 := "HP"
	brand2 := "Dell"
	brand3 := "Asus"

	brand := ""
	fmt.Println("Бренд который есть в магазине:")
	fmt.Scan(&brand)

	//result = brand == brand1 || brand == brand2
	//result = brand != brand3
	//fmt.Printf("%d > 18 = %t\n", age, result)

	result = (brand == brand1 || brand == brand2) && brand != brand3

	// 2 + 2 * 2 = 6
	// (2 + 2) * 2 = 8
	fmt.Println(result)
}
