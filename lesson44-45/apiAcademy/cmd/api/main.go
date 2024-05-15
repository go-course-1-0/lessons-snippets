package main

import (
	"apiAcademy/internal/database/models"
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

	handlers := Handlers{
		DB: db,
	}

	router := gin.Default()

	admins := router.Group("/admins")
	{
		admins.GET("/", handlers.GetAllAdmins)
		admins.POST("/", handlers.CreateAdmin)
		admins.GET("/:id", handlers.GetAllAdmins)
		admins.PUT("/:id", handlers.GetAllAdmins)
		admins.DELETE("/:id", handlers.GetAllAdmins)
	}

	teachers := router.Group("/teachers")
	{
		teachers.GET("/", handlers.GetAllAdmins)
		teachers.POST("/", handlers.GetAllAdmins)
		teachers.GET("/:id", handlers.GetAllAdmins)
		teachers.PUT("/:id", handlers.GetAllAdmins)
		teachers.DELETE("/:id", handlers.GetAllAdmins)
	}

	courses := router.Group("/courses")
	{
		courses.GET("/", handlers.GetAllAdmins)
		courses.POST("/", handlers.GetAllAdmins)
		courses.GET("/:id", handlers.GetAllAdmins)
		courses.PUT("/:id", handlers.GetAllAdmins)
		courses.DELETE("/:id", handlers.GetAllAdmins)
	}

	groups := router.Group("/groups")
	{
		groups.GET("/", handlers.GetAllAdmins)
		groups.POST("/", handlers.GetAllAdmins)
		groups.GET("/:id", handlers.GetAllAdmins)
		groups.PUT("/:id", handlers.GetAllAdmins)
		groups.DELETE("/:id", handlers.GetAllAdmins)
	}

	students := router.Group("/students")
	{
		students.GET("/", handlers.GetAllAdmins)
		students.POST("/", handlers.GetAllAdmins)
		students.GET("/:id", handlers.GetAllAdmins)
		students.PUT("/:id", handlers.GetAllAdmins)
		students.DELETE("/:id", handlers.GetAllAdmins)
	}

	lessons := router.Group("/lessons")
	{
		lessons.GET("/", handlers.GetAllAdmins)
		lessons.POST("/", handlers.GetAllAdmins)
		lessons.GET("/:id", handlers.GetAllAdmins)
		lessons.PUT("/:id", handlers.GetAllAdmins)
		lessons.DELETE("/:id", handlers.GetAllAdmins)
	}

	router.Run(":4000")
}

type Handlers struct {
	DB *gorm.DB
}

func (h *Handlers) GetAllAdmins(c *gin.Context) {
	var admins []models.Admin
	h.DB.Find(&admins)

	c.JSON(200, admins)
}

func (h *Handlers) CreateAdmin(c *gin.Context) {
	var admin models.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(422, "Validation error")
		return
	}

	h.DB.Create(&admin)

	c.JSON(201, admin)
}
