package db

import (
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Get(connStr string) (*gorm.DB, error) {
	println(connStr)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{}) //db, err := get(connStr)
	if err != nil {
		log.Fatal("Connect database failed!", err)
		return nil, err
	}
	log.Info("Connect Database successfully!!")
	return db, nil
}
