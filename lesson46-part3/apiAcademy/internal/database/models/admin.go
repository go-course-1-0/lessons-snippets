package models

import "time"

type Admin struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	FullName  string    `json:"fullName" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null;unique"`
	Password  string    `json:"-" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"default:now()"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"default:now()"`
}
