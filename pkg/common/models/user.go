package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"not null"`
	Email string `gorm:"not null"`
	PasswordHash string `gorm:"not null"`
	Sleeps []Sleep
}