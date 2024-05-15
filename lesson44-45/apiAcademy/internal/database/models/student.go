package models

import "time"

type Student struct {
	ID          int       `gorm:"primaryKey"`
	GroupID     int       `gorm:"not null;"`
	FullName    string    `gorm:"not null"`
	Phone       string    `gorm:"not null;unique"`
	DateOfBirth time.Time `gorm:"not null"`
	CreatedAt   time.Time `gorm:"default:now()"`
	UpdatedAt   time.Time `gorm:"default:now()"`

	// student belongs to group
	Group *Group
}
