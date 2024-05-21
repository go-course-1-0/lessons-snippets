package models

import "time"

type Teacher struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	FullName  string    `json:"fullName" gorm:"not null"`
	Subject   string    `json:"subject" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:now()"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:now()"`

	// teacher has many groups
	Groups []Group `json:"groups,omitempty"`
}
