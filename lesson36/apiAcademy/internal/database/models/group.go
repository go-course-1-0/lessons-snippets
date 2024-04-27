package models

import "time"

type Group struct {
	ID        int
	CourseID  int
	TeacherID int
	Title     string
	Start     time.Time
	Finish    time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
