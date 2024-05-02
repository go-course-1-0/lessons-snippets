package main

import (
	"apiAcademy/internal/examples"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost" // 127.0.0.1
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "academy_db"
)

func main() {
	// DSN - data source name
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("cannot connect to db:", err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("cannot ping db:", err.Error())
	}

	log.Println("Successfully connected!")

	//examples.Example1(db)
	//examples.Example2SqlInjection(db)

	//examples.StaticQuery(db)
	//examples.DynamicQuery(db)
	//examples.DynamicQueryPrepared(db)
	examples.DynamicQueryPreparedUsingStruct(db)
}
