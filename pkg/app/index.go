package app

import (
	"gorm.io/gorm"
	"grpc-echo/pkg/cache"
	"grpc-echo/pkg/config"
	"grpc-echo/pkg/db"
	"grpc-echo/pkg/models"
	"log"
)

// Application holds commonly used app wide data, for ease of DI
type Application struct {
	DB  *db.DB
	Cfg *config.Config
	RC  *cache.Client
}

// Get captures env vars, establishes DB connection and keeps/returns
// reference to both
func Get() (*Application, error) {
	cfg := config.Get()
	dbData, err := db.Get(cfg.GetDBConnStr())
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	errMr := dbData.Client.AutoMigrate(&models.Students{})
	if errMr != nil {
		log.Panic(errMr)
	}
	client, err := cache.InitRedisClient(cfg.GetRedisUrl(), cfg.GetRedisPassword(), cfg.GetRedisDB())
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	return &Application{
		DB:  dbData,
		Cfg: cfg,
		RC:  client,
	}, nil
}

func GetDB() (*gorm.DB, error) {
	rs, err := Get()
	if err != nil {
		return nil, err
	}
	return rs.DB.Client, nil

}

func GetCache() (*cache.Client, error) {
	rs, err := Get()
	if err != nil {
		return nil, err
	}
	return rs.RC, nil
}
