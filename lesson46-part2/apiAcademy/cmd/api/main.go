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
		teachers.GET("/", h.GetAllTeachers)
		teachers.POST("/", h.CreateTeacher)
		teachers.GET("/:id", h.GetOneTeacher)
		teachers.PUT("/:id", h.UpdateTeacher)
		teachers.DELETE("/:id", h.DeleteTeacher)
	}

	courses := router.Group("/courses")
	{
		courses.GET("/", h.GetAllCourses)
		courses.POST("/", h.CreateCourse)
		courses.GET("/:id", h.GetOneCourse)
		courses.PUT("/:id", h.UpdateCourse)
		courses.DELETE("/:id", h.DeleteCourse)
	}

	groups := router.Group("/groups")
	{
		groups.GET("/", h.GetAllGroups)
		groups.POST("/", h.CreateGroup)
		groups.GET("/:id", h.GetOneGroup)
		groups.PUT("/:id", h.UpdateGroup)
		groups.DELETE("/:id", h.DeleteGroup)
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
