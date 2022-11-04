package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)
type Sleep struct {
	gorm.Model
	DATE datatypes.Date `gorm:"not null"`
	SLEEP_START datatypes.Time `gorm:"not null"`
	SLEEP_END datatypes.Time `gorm:"not null"`
	UserID uint `gorm:"not null"`
}