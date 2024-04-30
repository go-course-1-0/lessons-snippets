package models

import "time"

type Teacher struct {
	ID        int
	FullName  string
	Subject   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
