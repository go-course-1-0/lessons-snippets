package models

import "time"

type Course struct {
	ID        int       `gorm:"primaryKey"`
	Title     string    `gorm:"not null"`
	Duration  int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt time.Time `gorm:"default:now()"`

	// course has many groups
	Groups []Group
}
