package main

import (
	"apiAcademy/internal/database/models"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	host     = "localhost" // 127.0.0.1
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "academy_db"
)

func main() {
	// 1. Self written SQL method:
	//db, err := database.UsualConnect(host, port, user, password, dbname)
	//if err != nil {
	//	log.Fatal("cannot connect to DB:", err.Error())
	//}
	//defer db.Close()

	//examples.Example1(db)
	//examples.Example2SqlInjection(db)

	//examples.StaticQuery(db)
	//examples.DynamicQuery(db)
	//examples.DynamicQueryPrepared(db)
	//examples.DynamicQueryPreparedUsingStruct(db)

	// 2. ORM connection method:
	// DSN - data source name
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot open gorm connection:", err.Error())
	}

	err = db.AutoMigrate(
		&models.Admin{},
		&models.Teacher{},
		&models.Course{},
		&models.Group{},
		&models.Student{},
		&models.Lesson{},
	)
	if err != nil {
		log.Fatal("cannot auto-migrate the DB:", err.Error())
	}

	fmt.Println("All good")
}
