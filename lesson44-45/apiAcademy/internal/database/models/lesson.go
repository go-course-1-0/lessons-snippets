package models

import "time"

type Lesson struct {
	ID        int       `gorm:"primaryKey"`
	GroupID   int       `gorm:"not null;"`
	DayOfWeek int       `gorm:"not null"`
	Time      time.Time `gorm:"default:null"`
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt time.Time `gorm:"default:now()"`

	// lesson belongs to group
	Group *Group
}
