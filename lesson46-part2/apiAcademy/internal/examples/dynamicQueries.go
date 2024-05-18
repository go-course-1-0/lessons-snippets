package examples

import (
	"apiAcademy/internal/database/models"
	"database/sql"
	"fmt"
	"log"
)

func StaticQuery(db *sql.DB) {
	var (
		id       int64
		fullName string
		email    string
		password string
	)

	fmt.Println("Creating new admin. Enter admin's data:")
	fmt.Print("Full name: ")
	fmt.Scan(&fullName)
	fmt.Print("Email: ")
	fmt.Scan(&email)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	query := fmt.Sprintf("insert into admins (full_name, email, password) values ('%s', '%s', '%s')", fullName, email, password)
	fmt.Println(query)
	result, err := db.Exec(query)
	if err != nil {
		log.Fatal("cannot insert into admins (static):", err.Error())
	}

	id, err = result.LastInsertId()
	if err != nil {
		log.Fatal("cannot get LastInsertId (static):", err.Error())
	}

	fmt.Println("Admin created successfully!")
	fmt.Println("ID:", id)
	fmt.Println("Full Name:", fullName)
	fmt.Println("Email:", email)
	fmt.Println("Password:", password)
}

func DynamicQuery(db *sql.DB) {
	var (
		id       int64
		fullName string
		email    string
		password string
	)

	fmt.Println("Creating new admin. Enter admin's data:")
	fmt.Print("Full name: ")
	fmt.Scan(&fullName)
	fmt.Print("Email: ")
	fmt.Scan(&email)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	query := "insert into admins (full_name, email, password) values ($1, $2, $3)"
	fmt.Println(query)
	result, err := db.Exec(query, fullName, email, password)
	if err != nil {
		log.Fatal("cannot insert into admins (dynamic):", err.Error())
	}

	id, err = result.LastInsertId()
	if err != nil {
		log.Fatal("cannot get LastInsertId (dynamic):", err.Error())
	}

	fmt.Println("Admin created successfully!")
	fmt.Println("ID:", id)
	fmt.Println("Full Name:", fullName)
	fmt.Println("Email:", email)
	fmt.Println("Password:", password)
}

func DynamicQueryPrepared(db *sql.DB) {
	var (
		id       int64
		fullName string
		email    string
		password string
	)

	fmt.Println("Creating new admin. Enter admin's data:")
	fmt.Print("Full name: ")
	fmt.Scan(&fullName)
	fmt.Print("Email: ")
	fmt.Scan(&email)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	statement, err := db.Prepare("insert into admins (full_name, email, password) values ($1, $2, $3)")
	if err != nil {
		log.Fatal("cannot prepare statement:", err.Error())
	}

	defer statement.Close()

	_, err = statement.Exec(fullName, email, password)
	if err != nil {
		log.Fatal("cannot insert into admins (dynamic prepared):", err.Error())
	}

	fmt.Println("Admin created successfully!")
	fmt.Println("ID:", id)
	fmt.Println("Full Name:", fullName)
	fmt.Println("Email:", email)
	fmt.Println("Password:", password)
}

func DynamicQueryPreparedUsingStruct(db *sql.DB) {
	var admin models.Admin

	fmt.Println("Creating new admin. Enter admin's data:")
	fmt.Print("Full name: ")
	fmt.Scan(&admin.FullName)
	fmt.Print("Email: ")
	fmt.Scan(&admin.Email)
	fmt.Print("Password: ")
	fmt.Scan(&admin.Password)

	statement, err := db.Prepare("insert into admins (full_name, email, password) values ($1, $2, $3)")
	if err != nil {
		log.Fatal("cannot prepare statement:", err.Error())
	}

	defer statement.Close()

	_, err = statement.Exec(admin.FullName, admin.Email, admin.Password)
	if err != nil {
		log.Fatal("cannot insert into admins (dynamic prepared with struct):", err.Error())
	}

	fmt.Println("Admin created successfully!")
	fmt.Println("Full Name:", admin.FullName)
	fmt.Println("Email:", admin.Email)
	fmt.Println("Password:", admin.Password)
}
