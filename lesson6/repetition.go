package main

import "fmt"

func main() {
	var color string
	fmt.Println("What is current color of the traffic light?")
	fmt.Scan(&color)

	switch color {
	case "red":
		fmt.Println("Stop")
		fallthrough
	case "yellow":
		fmt.Println("Ready")
		fallthrough
	case "green":
		fmt.Println("Go")
	default:
		fmt.Println("Wrong color")
	}

	//switch {
	//case color == "red":
	//	fmt.Println("Stop")
	//	time.Sleep(time.Second * 5)
	//	fallthrough
	//case color == "yellow":
	//	fmt.Println("Ready")
	//	time.Sleep(time.Second * 2)
	//	fallthrough
	//case color == "green":
	//	fmt.Println("Go")
	//	time.Sleep(time.Second * 5)
	//default:
	//	fmt.Println("Wrong color")
	//}
}
