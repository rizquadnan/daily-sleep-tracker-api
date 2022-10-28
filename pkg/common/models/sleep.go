package models

import "gorm.io/gorm"

type Sleep struct {
	gorm.Model
	DATE string `json:"date"`
	SLEEP_START string `json:"sleep_start"`
	SLEEP_END string `json:"sleep_end"`
}