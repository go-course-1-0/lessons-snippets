package main

import (
	"studentsListManagement/internal/models"
)

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

func main() {
	app := models.NewApplication()

	if err := app.Run(); err != nil {
		app.ErrorLogger.Error("cannot run application:", err.Error())
	}
}
