package database

import (
	"database/sql"
	"fmt"
)

func UsualConnect(host, port, user, password, dbname string) (*sql.DB, error) {
	// DSN - data source name
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
