package app

import (
	"gorm.io/gorm"
	"grpc-echo/pkg/cache"
	"grpc-echo/pkg/config"
	"grpc-echo/pkg/db"
	"log"
	"os"
)

// Application holds commonly used app wide data, for ease of DI
type Application struct {
	DB  *gorm.DB
	RC  *cache.Client
}

func Init() error  {
	cfg := config.Get()
	_, err := db.Get(cfg.GetDBConnStr())
	if err != nil {
		log.Panic(err)
		return err
	}
	_, errRd := cache.InitRedisClient(cfg.GetRedisUrl(), cfg.GetRedisPassword(), cfg.GetRedisDB())
	if errRd != nil {
		log.Panic(errRd)
		return errRd
	}
	return nil

}
func GetAPIPort() string {
	cfg := config.Get()
	return cfg.GetAPIPort()
}
func Get() (*Application, error) {
	cfg := config.Get()
	dbData, err := db.Get(cfg.GetDBConnStr())
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	client, err := cache.InitRedisClient(cfg.GetRedisUrl(), cfg.GetRedisPassword(), cfg.GetRedisDB())
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	return &Application{
		DB:  dbData,
		RC:  client,
	}, nil
}

func GetDB() (*gorm.DB, error) {

	//cfg := config.Get()
	//dbData, err := db.Get(cfg.GetDBConnStr())
	//const DBconn = "postgres://test:123456@127.0.0.1:5432/test?sslmode=disable"
	dbData, err := db.Get(os.Getenv("POSTGRES_URL"))
	dbData.Migrator()
	if err != nil {
		return nil, err
	}
	return dbData, nil
}

func GetCache() (*cache.Client, error) {
	cfg := config.Get()
	client, err := cache.InitRedisClient(cfg.GetRedisUrl(), cfg.GetRedisPassword(), cfg.GetRedisDB())
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	return client, nil
}
