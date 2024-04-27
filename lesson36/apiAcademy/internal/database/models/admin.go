package models

import "time"

type Admin struct {
	ID        int
	FullName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
