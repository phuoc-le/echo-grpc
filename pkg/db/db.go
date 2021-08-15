package db

import (
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Client *gorm.DB
}

func Get(connStr string) (*DB, error) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{}) //db, err := get(connStr)
	if err != nil {
		log.Fatal("Connect database failed!", err)
		return nil, err
	}
	log.Info("Connect Database successfully!!")
	return &DB{
		Client: db,
	}, nil
}
