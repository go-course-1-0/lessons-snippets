package models

import "time"

type Course struct {
	ID        int
	Title     string
	Duration  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
