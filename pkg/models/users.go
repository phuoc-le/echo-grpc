package models

import (
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	UserName string  `gorm:"not null"`
	FirstName string
	LastName string
	Password string
	groups string
	user_permissions string
	is_staff bool
	isActive bool
	is_superuser bool
	last_login pgtype.Date
	Email string `gorm:"unique,not null,"`
	Phone string
}
