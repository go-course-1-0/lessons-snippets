package models

import "time"

type Lesson struct {
	ID          int
	GroupID     int
	DayOfWeek   int
	Time        time.Time
	DateOfBirth time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
