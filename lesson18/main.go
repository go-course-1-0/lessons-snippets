package main

import "fmt"

type Number int
type Age int
type Name string
type Surname string
type Weight float64
type Height float64

type Person struct {
	Name     string
	Surname  string
	Age      int
	Weight   float64
	Height   float64
	Passport int
}

type Student struct {
	Person
	University string
	Course     int
	Group      string
	Computer   Laptop
	Books      []Book
}

func (s Student) Study() {
	fmt.Println(s.Name, "is Studying")
}

type Book struct {
	Title string
}

type Laptop struct {
	Brand      string
	ScreenSize int
}

type Teacher struct {
	Person
	TeachingTopic string
}

func main() {
	var name Name = "Siyovush"
	var age Age = 24
	var weight Weight = 70.5
	var height Height = 170.8

	fmt.Println(name, age, weight, height)

	var person Person

	person.Name = "Faridun"
	person.Age = 29
	person.Weight = 85
	person.Height = 179.5

	fmt.Println(person)

	anotherPerson := Person{
		Name:   "Ubaidullo",
		Age:    18,
		Weight: 65,
		Height: 162,
	}

	fmt.Println(anotherPerson)

	var pName = &name
	fmt.Println(pName, *pName)

	pPerson := &person
	fmt.Println(pPerson, *pPerson)

	name += " Shorakhimov"
	*pName += " Makhsudovich"
	fmt.Println(name)

	fmt.Println(person.Name)
	fmt.Println((*pPerson).Name)
	fmt.Println(pPerson.Name)

	//otherPerson := Person{"Siyovush", "sho", 23, 56, 180}
	//fmt.Println(otherPerson)

	st := Student{
		Person: Person{
			Name:    "Siyovush",
			Surname: "Ivanov",
			Age:     25,
			Weight:  80,
			Height:  180,
		},
		University: "ITMO",
		Course:     3,
		Group:      "9876A",
	}

	fmt.Println(st)

	st1 := Student{
		Person: Person{
			Name:    "Ubaidullo",
			Surname: "Ivanov 1",
			Age:     25,
			Weight:  80,
			Height:  180,
		},
		University: "ITMO",
		Course:     3,
		Group:      "9876A",
		Computer: Laptop{
			Brand:      "HP",
			ScreenSize: 16,
		},
	}

	fmt.Println(st1)

	fmt.Println(st1.University)
	fmt.Println(st1.Person)
	fmt.Println(st1.Computer)

	fmt.Println(st1.Computer.Brand)
	fmt.Println(st1.Person.Name)
	fmt.Println(st1.Age)
	fmt.Println(st1.Computer.Brand)

	st1.Study()
	st.Study()
}
