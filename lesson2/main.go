package main

import "fmt"

func main() {
	// numbers
	// string
	// bool

	var b bool
	fmt.Println(b)
	b = true
	fmt.Println(b)

	fmt.Println("\nAnother type\n")

	var s string
	fmt.Println(s)
	s = "Go Course 1.1"
	fmt.Println(s)

	fmt.Println("\nAnother type\n")

	var a int
	fmt.Println(a)
	a = 10
	fmt.Println(a)

	fmt.Println("\nAnother type\n")

	var r rune
	fmt.Println(r)
	r = 65
	fmt.Println(string(r))

	fmt.Println("\nAnother type\n")

	var f float32
	fmt.Println(f)
	f = 10.9
	fmt.Printf("Siyovush has %.2f somoni\n", f)
	fmt.Print("Siyovush has ", f, " somoni\n")

	fmt.Println("\nAnother type\n")

	fmt.Println(float32(a) + f)
	fmt.Println(a + int(f))

	var big uint64
	big = 999_999_999_999
	fmt.Println(big)

	var small uint16
	fmt.Println(small)
	small = uint16(big)
	fmt.Println(small)

	// Run the program

	// Ask the user's name
	// Ask the user's year of birth
	// Calculate his age
	// Print: Name you are ? years old

	// Ask how much he earns per month
	// Ask how much he wants to save money
	// Calculate the amount of months needed to save money for him
}
