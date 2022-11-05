package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"not null" json:"name"`
	Email string `gorm:"not null" json:"email"`
	PasswordHash string `gorm:"not null" json:"passwordHash"`
	Sleeps []Sleep `json:"sleeps"`
}