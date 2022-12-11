package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)
type Sleep struct {
	gorm.Model
	DATE datatypes.Date `gorm:"not null" json:"date"`
	SLEEP_START datatypes.Time `gorm:"not null" json:"sleepStart"`
	SLEEP_END datatypes.Time `gorm:"not null" json:"sleepEnd"`
	SLEEP_DURATION int `gorm:"not null" json:"sleepDuration"`
	UserID uint `gorm:"not null" json:"userId"`
}