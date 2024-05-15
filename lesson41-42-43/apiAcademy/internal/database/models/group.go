package models

import "time"

type Group struct {
	ID        int       `gorm:"primaryKey"`
	CourseID  int       `gorm:"not null;"`
	TeacherID int       `gorm:"not null;"`
	Title     string    `gorm:"not null"`
	Start     time.Time `gorm:"default:null"`
	Finish    time.Time `gorm:"default:null"`
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt time.Time `gorm:"default:now()"`

	// group belongs to teacher
	Teacher *Teacher
	// group belongs to course
	Course *Course

	// group has many students
	Students []Student
	// group has many lessons
	Lessons []Lesson
}
