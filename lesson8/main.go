package main

import "fmt"

func main() {
	//var year int
	//fmt.Println("Год рождения:")
	//fmt.Scan(&year)
	//
	//age := calculateAge(year)
	//fmt.Println("Возраст:")
	//fmt.Println(age)
	//
	//fmt.Println("Может получить водительское удостоверение:")
	//fmt.Println(isAbleToGetDriverLicense(age))

	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println("Периметр:")
	fmt.Println(perimeterOfRectangle(a, b))

	fmt.Println("Периметр и площадь:")
	fmt.Println(rectangleCharacteristics(a, b))

	var perimeter, square int

	perimeter, square = rectangleCharacteristics(a, b)
	fmt.Println(perimeter, square)

	_, square = rectangleCharacteristics(a, b)
	fmt.Println(perimeter, square)
}

func calculateAge(yearOfBirth int) int {
	return 2024 - yearOfBirth
}

func isAbleToGetDriverLicense(age int) bool {
	return age >= 18
}

func perimeterOfRectangle(a, b int) int {
	return 2 * (a + b)
}

func rectangleCharacteristics(a, b int) (p int, s int) {
	return 2 * (a + b), a * b
}
