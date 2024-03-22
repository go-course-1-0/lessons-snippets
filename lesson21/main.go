package main

import (
	"fmt"
	"itrun-lectures/lesson21/generalFunctions"
	"itrun-lectures/lesson21/geometry"
)

func main() {
	generalFunctions.SayHello()
	generalFunctions.SayGoodBye()
	//generalFunctions.sayNothing()

	fmt.Println(geometry.PerimeterOfSquare(10))
	fmt.Println(geometry.PerimeterOfRectangle(10, 20))
	fmt.Println(geometry.PerimeterOfRectangleTriangle(10, 20, 30))
}
