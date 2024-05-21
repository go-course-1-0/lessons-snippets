package models

import "time"

type Course struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Duration  int       `json:"duration" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:now()"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:now()"`

	// course has many groups
	Groups []Group `json:"groups,omitempty"`
}
