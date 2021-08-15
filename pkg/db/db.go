package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Client *gorm.DB
}

func Get(connStr string) (*DB, error) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{}) //db, err := get(connStr)
	if err != nil {
		return nil, err
	}

	return &DB{
		Client: db,
	}, nil
}
