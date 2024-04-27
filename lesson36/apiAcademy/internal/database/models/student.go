package models

import "time"

type Student struct {
	ID          int
	GroupID     int
	FullName    string
	Phone       string
	DateOfBirth time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
