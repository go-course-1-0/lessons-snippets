package main

import (
	"apiAcademy/internal/database/models"
	"apiAcademy/internal/handlers"
	"apiAcademy/internal/helpers"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	DBHost     = "localhost" // 127.0.0.1
	DBPort     = 5432
	DBUser     = "postgres"
	DBPassword = "postgres"
	DBName     = "academy_db"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Dushanbe", DBHost, DBPort, DBUser, DBPassword, DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("successfully connected to the DB")

	if err = db.AutoMigrate(
		&models.Admin{},
		&models.Teacher{},
		&models.Course{},
		&models.Group{},
		&models.Student{},
		&models.Lesson{},
	); err != nil {
		log.Fatal("cannot migrate tables:", err.Error())
	}

	helpers.SetValidatorEngineToUseJSONTags()

	h := handlers.Handlers{
		DB: db,
	}

	router := gin.Default()

	admins := router.Group("/admins")
	{
		admins.GET("/", h.GetAllAdmins)      // #done #withoutPagination
		admins.POST("/", h.CreateAdmin)      // #done
		admins.GET("/:id", h.GetOneAdmin)    // #done
		admins.PUT("/:id", h.UpdateAdmin)    // #done
		admins.DELETE("/:id", h.DeleteAdmin) // #done
	}

	teachers := router.Group("/teachers")
	{
		teachers.GET("/", h.GetAllTeachers)      // #done #withoutPagination
		teachers.POST("/", h.CreateTeacher)      // #done
		teachers.GET("/:id", h.GetOneTeacher)    // #done
		teachers.PUT("/:id", h.UpdateTeacher)    // #done
		teachers.DELETE("/:id", h.DeleteTeacher) // #done
	}

	courses := router.Group("/courses")
	{
		courses.GET("/", h.GetAllCourses)      // #done #withoutPagination
		courses.POST("/", h.CreateCourse)      // #done
		courses.GET("/:id", h.GetOneCourse)    // #done
		courses.PUT("/:id", h.UpdateCourse)    // #done
		courses.DELETE("/:id", h.DeleteCourse) // #done
	}

	groups := router.Group("/groups")
	{
		groups.GET("/", h.GetAllGroups)      // #done #withoutPagination
		groups.POST("/", h.CreateGroup)      // #done #withoutForeignKeysValidation #withoutDateValidation
		groups.GET("/:id", h.GetOneGroup)    // #done
		groups.PUT("/:id", h.UpdateGroup)    // #done #withoutForeignKeysValidation #withoutDateValidation
		groups.DELETE("/:id", h.DeleteGroup) // #done
	}

	students := router.Group("/students")
	{
		students.GET("/", h.GetAllStudents)
		students.POST("/", h.CreateStudent)
		students.GET("/:id", h.GetOneStudent)
		students.PUT("/:id", h.UpdateStudent)
		students.DELETE("/:id", h.DeleteStudent)
	}

	lessons := router.Group("/lessons")
	{
		lessons.GET("/", h.GetAllLessons)
		lessons.POST("/", h.CreateLesson)
		lessons.GET("/:id", h.GetOneLesson)
		lessons.PUT("/:id", h.UpdateLesson)
		lessons.DELETE("/:id", h.DeleteLesson)
	}

	router.Run(":4000")
}
