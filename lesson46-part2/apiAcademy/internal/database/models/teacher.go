package models

import "time"

type Teacher struct {
	ID        int       `gorm:"primaryKey"`
	FullName  string    `gorm:"not null"`
	Subject   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt time.Time `gorm:"default:now()"`

	// teacher has many groups
	Groups []Group
}
