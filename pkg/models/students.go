package models

import "gorm.io/gorm"

type Students struct {
	gorm.Model
	Name string  `gorm:"not null"`
	Email string `gorm:"unique,not null,"`
	Phone string
}
