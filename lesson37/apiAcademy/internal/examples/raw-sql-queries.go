package examples

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func Example1(db *sql.DB) {
	rows, err := db.Query("SELECT count(*) FROM admins")
	if err != nil {
		log.Fatal("cannot execute query:", err.Error())
	}

	count := 0
	cnt := 0
	rows.Next()
	err = rows.Scan(&count)
	if err != nil {
		log.Fatal("cannot scan from row:", err.Error())
	}
	fmt.Println(count)
	cnt++
	fmt.Println(cnt, "rows")

	fmt.Println("The second query:\n")

	rows, err = db.Query("SELECT * FROM admins")
	if err != nil {
		log.Fatal("cannot execute query:", err.Error())
	}

	var (
		id        int
		fullName  string
		email     string
		pass      string
		createdAt time.Time
		updatedAt time.Time
	)

	cnt = 0
	for rows.Next() {
		err := rows.Scan(&id, &fullName, &email, &pass, &createdAt, &updatedAt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, fullName, email, pass, createdAt, updatedAt)
		cnt++
	}

	fmt.Println(cnt, "rows")
}

func Example2SqlInjection(db *sql.DB) {
	//ask:
	fmt.Println("Enter the table name to display:")
	var tableName string
	//fmt.Scanln(&tableName)

	tableName = "teachers; drop table students; drop table lessons; drop table groups;"
	//whiteList := []string{"admins", "teachers", "courses"}

	//if !slices.Contains(whiteList, tableName) {
	//	fmt.Println("Wrong table name")
	//goto ask
	//}

	query := fmt.Sprintf("SELECT id, created_at, updated_at FROM %s", tableName)

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("cannot execute query:", err.Error())
	}

	//switch tableName {
	//case "admins":
	//case "teachers":
	//case "courses":
	//case "groups":
	//case "students":
	//case "lessons":
	//}

	var (
		id        int
		createdAt time.Time
		updatedAt time.Time
	)

	cnt := 0
	for rows.Next() {
		err := rows.Scan(&id, &createdAt, &updatedAt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, createdAt, updatedAt)
		cnt++
	}

	fmt.Println(cnt, "rows")
}

// ORM - object relational mapping
// gORM //
