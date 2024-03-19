package main

import "fmt"

type Writer interface {
	Write()
}

type Pen struct {
	Color string
}

func (p Pen) Write() {
	fmt.Println("writing with pen")
}

type Pencil struct {
	Length float64
}

func (p Pencil) Write() {
	fmt.Println("writing with pencil")
}

func writeWithPen(p Pen) {
	fmt.Printf("I am writing this with %T of color %s\n", p, p.Color)
}

func writeWithPencil(p Pencil) {
	fmt.Printf("I am writing this with %T with length %f\n", p, p.Length)
}

func write(i Writer) {
	fmt.Printf("I am writing this with %T:\n", i)
	//i.Write()
	pen, ok := i.(Pen)
	if ok {
		pen.Write()
	}

	fmt.Println()
}

func main() {
	//fmt.Print("hello\n")
	//pen := Pen{Color: "black"}
	//writeWithPen(pen)
	//
	//pencil := Pencil{Length: 10}
	//writeWithPencil(pencil)
	//
	//write(pen)
	//write(pencil)

	//u := User{
	//	Name:     "Siyovush",
	//	Email:    "123@123.com",
	//	Password: "password",
	//}
	//SaveToStorage(u)
	//SaveToStorage(123)
	//SaveToStorage(123.57)
	//
	//c := Course{
	//	Title:    "Go Programming Language",
	//	Duration: 4,
	//}
	//SaveToStorage(c)
	//SaveToStorage(true)
	//SaveToStorage(map[int]string{1: "Hello", 2: "World"})
	//
	//n := Notification{
	//	Source: "Stepik",
	//	Text:   "Вас давно тут не было. Заходите!",
	//}
	//SaveToStorage(n)

	var a = 45

	fmt.Printf("%v - %T\n", a, a)
	fmt.Println()

	var i interface{}
	fmt.Printf("%v - %T\n", i, i)
	fmt.Println()

	i = a
	fmt.Printf("%v - %T\n", i, i)
	fmt.Println()

	i = "a"
	fmt.Printf("%v - %T\n", i, i)
	fmt.Println()

	i = true
	fmt.Printf("%v - %T\n", i, i)
	fmt.Println()

}

type User struct {
	Name     string
	Email    string
	Password string
}

type Course struct {
	Title    string
	Duration int
}

type Notification struct {
	Source string
	Text   string
}

func SaveToStorage(i interface{}) {
	fmt.Printf("Interface variable: %v - %T\n", i, i)

	//user, ok := i.(User)
	//if ok {
	//	fmt.Printf("Storing %s to users table of the DB\n", user.Name)
	//}
	//
	//course, ok := i.(Course)
	//if ok {
	//	fmt.Printf("Storing %s course with a duration of %d months to courses.txt file\n", course.Title, course.Duration)
	//}
	//
	//notification, ok := i.(Notification)
	//if ok {
	//	fmt.Printf("Storing notification from %s to notifications map\n", notification.Source)
	//}

	if user, ok := i.(User); ok {
		fmt.Printf("Storing %s to users table of the DB\n", user.Name)
	} else if course, ok := i.(Course); ok {
		fmt.Printf("Storing %s course with a duration of %d months to courses.txt file\n", course.Title, course.Duration)
	} else if notification, ok := i.(Notification); ok {
		fmt.Printf("Storing notification from %s to notifications map\n", notification.Source)
	}

	//switch t {
	//case "user":
	//case "course":
	//case "notification":
	//}

	//fmt.Printf("User variable: %v - %T\n", user, user)
	//fmt.Println("Assertion is successful:", ok)
	fmt.Println()
}
