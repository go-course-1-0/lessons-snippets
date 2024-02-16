package main

import "fmt"

func main() {
	//var mainTank, additionalTank int
	//
	//fmt.Println("Liters of fuel in the main tank:")
	//fmt.Scan(&mainTank)
	//
	//fmt.Println("Liters of fuel in the additional tank:")
	//fmt.Scan(&additionalTank)
	//
	//// your operations
	//distance := 0
	//liters := 0
	//for mainTank > 0 {
	//
	//	mainTank--
	//	liters++
	//	distance += 10
	//
	//	if liters == 5 {
	//		if additionalTank >= 1 {
	//			additionalTank--
	//			mainTank++
	//		}
	//		liters = 0
	//	}
	//}
	//
	//fmt.Println(distance)

	var mainTank, additionalTank, distanceTraveled int

	fmt.Println("Liters of the fuel in the main tank:")
	fmt.Scan(&mainTank)

	fmt.Println("Liters of fuel in the additional tank:")
	fmt.Scan(&additionalTank)

	for i := mainTank; i != 0; i-- {
		if i%5 == 0 {
			distanceTraveled += 10
		}
		distanceTraveled += 10
	}
	fmt.Println(distanceTraveled)
}
