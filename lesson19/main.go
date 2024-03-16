package main

import "fmt"

// Application
// Student list management application

// #done To exit
// #done Show the list of available commands						list-commands
// #done Show the list of students									list-students
// #done Add new student											add-student
// 		check number of students limit
//		check uniqueness of students
// #done Show the specific student // Find student by some identifier		show-student
// Update the specific student										update-student
// Delete the specific student										delete-student

var app application

const MaxNumberOfStudents = 10

func main() {
	app.seedStudentsList()

	for {
		var command string
		fmt.Println("Hello! Enter your command:")
		fmt.Scan(&command)

		if command == "exit" {
			return
		} else if command == "list-commands" {
			app.showAvailableCommands()
		} else if command == "list-students" {
			app.showStudentsList()
		} else if command == "add-student" {
			app.addStudent()
		} else if command == "show-student" {
			app.showStudent()
		}
	}
}

type student struct {
	FullName string
	Phone    string
}

type application struct {
	students []student
}

func (a *application) showAvailableCommands() {
	commands := []string{
		"list-commands",
		"list-students",
		"add-student",
		"show-student",
		"update-student",
		"delete-student",
	}

	fmt.Println()
	fmt.Println("I can run these commands:")
	for index, command := range commands {
		fmt.Printf("%d. %s\n", index+1, command)
	}
	fmt.Println()
}

func (a *application) showStudentsList() {
	fmt.Println()
	if len(a.students) == 0 {
		fmt.Println("There are no students in our academy.")
		fmt.Println()

		return
	}

	fmt.Println("These are the students of our academy:")
	for index, s := range a.students {
		fmt.Printf("%d. Name: %s;\tPhone: %s\n", index+1, s.FullName, s.Phone)
	}
	fmt.Println()
}

func (a *application) addStudent() {
	// Create new empty student object
	// Ask new student's info and fill the new object
	// Add this new student object to current student's list

	var s student
	fmt.Println()
	fmt.Println("Adding new student... Enter his info:")

	// check length

	fmt.Print("\tFull name:\t")
	fmt.Scan(&s.FullName)
	fmt.Print("\tPhone number:\t")
	fmt.Scan(&s.Phone)

	// check phone uniqueness

	//modifiedList := append(a.students, s)
	//a.students = modifiedList

	a.students = append(a.students, s)

	fmt.Printf("All good! %s has been added to the list successfully\n", s.FullName)
	fmt.Println()
}

func (a application) showStudent() {
	fmt.Println()
	if len(a.students) == 0 {
		fmt.Println("There are no students in our academy.")
		fmt.Println()

		return
	}

	fmt.Println("Enter student's ID:")
	var id int
	fmt.Scan(&id)

	id--

	if id < 0 || id >= len(a.students) {
		fmt.Printf("Wrong student ID! Enter a number from %d to %d\n", 1, len(a.students))
		return
	}

	fmt.Printf("%d. Name: %s;\tPhone: %s\n", id+1, a.students[id].FullName, a.students[id].Phone)
	fmt.Println()
}

func (a *application) seedStudentsList() {
	a.students = []student{
		{
			FullName: "Faridun",
			Phone:    "123456",
		},
		{
			FullName: "Farhod",
			Phone:    "234567",
		},
		{
			FullName: "Shokirjon",
			Phone:    "45678",
		},
		{
			FullName: "Muhammadali",
			Phone:    "897654",
		},
		{
			FullName: "Ubaidullo",
			Phone:    "8967452",
		},
	}
}
