package models

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
)

var ErrIncorrectCommand = errors.New("введена неправильная команда")
var ErrIncorrectPhoneNumber = errors.New("некорректный номер телефона")
var ErrDuplicatedPhoneNumber = errors.New("дублированный номер телефона")
var ErrStudentNotFound = errors.New("студент не найден")

type Application struct {
	students    []student
	InfoLogger  *slog.Logger
	ErrorLogger *slog.Logger
}

func NewApplication() Application {
	var app Application
	app.InfoLogger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app.ErrorLogger = slog.New(slog.NewJSONHandler(os.Stderr, nil))
	app.SeedStudentsList()

	return app
}

func (a *Application) Run() error {
	for {
		var command string
		fmt.Println("Hello! Enter your command:")
		fmt.Scan(&command)

		if command == "exit" {
			a.InfoLogger.Info("выполняется команда exit")
			return nil
		} else if command == "list-commands" {
			a.InfoLogger.Info("выполняется команда list-commands")
			a.ShowAvailableCommands()
		} else if command == "list-students" {
			a.InfoLogger.Info("выполняется команда list-students")
			a.ShowStudentsList()
		} else if command == "add-student" {
			a.InfoLogger.Info("выполняется команда add-student")
			a.AddStudent()
		} else if command == "show-student" {
			a.InfoLogger.Info("выполняется команда show-student")
			a.ShowStudent()
		} else {
			a.InfoLogger.Info(ErrIncorrectCommand.Error(), "command", command)
		}
	}
}

const MaxNumberOfStudents = 10

func (a *Application) ShowAvailableCommands() {
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

func (a *Application) ShowStudentsList() {
	fmt.Println()
	if len(a.students) == 0 {
		fmt.Println("There are no students in our academy.")
		fmt.Println()

		a.InfoLogger.Info("попытка просмотра пустого списка студентов")
		return
	}

	fmt.Println("These are the students of our academy:")
	for index, s := range a.students {
		fmt.Printf("%d. Name: %s;\tPhone: %s\n", index+1, s.FullName, s.Phone)
	}
	fmt.Println()
}

func (a *Application) AddStudent() {
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

	if len(s.Phone) != 9 {
		fmt.Println("Некорректный номер телефона!")
		a.ErrorLogger.Error(ErrIncorrectPhoneNumber.Error(), "phone", s.Phone)
		return
	}

	// check phone uniqueness

	var exists bool
	for _, student := range a.students {
		if student.Phone == s.Phone {
			exists = true
			break
		}
	}

	if exists {
		fmt.Println("Студент с таким номером телефона уже существует!")
		a.ErrorLogger.Error(ErrDuplicatedPhoneNumber.Error(), "phone", s.Phone)
		return
	}

	//modifiedList := append(a.students, s)
	//a.students = modifiedList

	a.students = append(a.students, s)

	fmt.Printf("All good! %s has been added to the list successfully\n", s.FullName)
	fmt.Println()
}

func (a *Application) ShowStudent() {
	fmt.Println()
	if len(a.students) == 0 {
		fmt.Println("There are no students in our academy.")
		fmt.Println()
		a.InfoLogger.Info("попытка просмотра пустого списка студентов")
		return
	}

	fmt.Println("Enter student's ID:")
	var id int
	fmt.Scan(&id)

	id--

	if id < 0 || id >= len(a.students) {
		fmt.Printf("Wrong student ID! Enter a number from %d to %d\n", 1, len(a.students))
		a.ErrorLogger.Error(ErrStudentNotFound.Error(), "studentID", id)
		return
	}

	fmt.Printf("%d. Name: %s;\tPhone: %s\n", id+1, a.students[id].FullName, a.students[id].Phone)
	fmt.Println()
}

func (a *Application) SeedStudentsList() {
	a.students = []student{
		{
			FullName: "Faridun",
			Phone:    "987654321",
		},
		{
			FullName: "Farhod",
			Phone:    "987654322",
		},
		{
			FullName: "Shokirjon",
			Phone:    "987654323",
		},
		{
			FullName: "Muhammadali",
			Phone:    "987654324",
		},
		{
			FullName: "Ubaidullo",
			Phone:    "8967452",
		},
	}
}
