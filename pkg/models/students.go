package models

import "gorm.io/gorm"

type Students struct {
	gorm.Model
	Id   string `json:"id"`
	Name string `json:"name"`
}
