package models

import "time"

type Group struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	CourseID  int       `json:"courseID" gorm:"not null;"`
	TeacherID int       `json:"teacherID" gorm:"not null;"`
	Title     string    `json:"title" gorm:"not null"`
	Start     time.Time `json:"start" gorm:"default:null"`
	Finish    time.Time `json:"finish" gorm:"default:null"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:now()"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:now()"`

	// group belongs to course
	Course *Course `json:"course,omitempty"`
	// group belongs to teacher
	Teacher *Teacher `json:"teacher,omitempty"`

	// group has many students
	Students []Student `json:"students,omitempty"`
	// group has many lessons
	Lessons []Lesson `json:"lessons,omitempty"`
}
