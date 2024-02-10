package main

import "fmt"

func main() {
	var name string
	fmt.Println("What is your name?")
	fmt.Print("My name is ")
	fmt.Scan(&name)

	switch name {
	case "Siyovush":
		fmt.Println("Hello teacher!")
	case "Farhod":
		fmt.Println("Hello Mr. Farhod")
	case "Muhammad":
		fmt.Println("Hello dear student")
	case "Komron":
		fmt.Println("Hello dear student!")
	case "Faridun", "Ubaidullo", "Shukur":
		fmt.Println("Hello dear student!")
	case "Shokir":
		fallthrough
	case "Abubakr":
		fallthrough
	case "Alijon":
		fmt.Println("Hello dear student!")
	default:
		fmt.Println("I don't know who are you")
	}

	age := 0

	fmt.Scan(&age)

	switch {
	case age < 18:
		fmt.Println("Sorry you cannot have a driving license")
	case name == "Siyovush":
		fmt.Println("Sorry you cannot have a driving license, because you are blocked!")
	default:
		fmt.Println("You can have a driving license")
	}
}
