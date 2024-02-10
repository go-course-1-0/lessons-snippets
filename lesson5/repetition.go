package main

import "fmt"

func main() {
	var name string
	fmt.Println("What is your name?")
	fmt.Print("My name is ")
	fmt.Scan(&name)

	if name == "Siyovush" {
		fmt.Println("Hello teacher!")
	} else if name == "Farhod" {
		fmt.Println(" Hello Mr. Farhod")
	} else if name == "Muhammad" {
		fmt.Println("Hello dear student")
	} else if name == "komron" {
		fmt.Println("Hello dear student!")
	} else if name == "Faridun" || name == "Ubaidullo" || name == "Shukur" {
		fmt.Println("Hello dear student")
	} else if name == "Olim" || name == "Emom" {
		fmt.Println("You have to pay for the lessons")
	} else {
		fmt.Println("I don't know who are you")
	}
}
