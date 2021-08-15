package storage

import (
	"grpc-echo/pkg/cache"
	"grpc-echo/pkg/config"
	"grpc-echo/pkg/db"
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

	db, err := db.Get(cfg.GetDBConnStr())
	if err != nil {
		return nil, err
	}
	client, err := cache.InitRedisClient(cfg.GetRedisUrl(), cfg.GetRedisPassword(), cfg.GetRedisDB())
	if err != nil {
		return nil, err
	}
	return &Application{
		DB:  db,
		Cfg: cfg,
		RC:  client,
	}, nil
}

func GetDB() (*db.DB, error) {
	rs, err := Get()
	if err != nil {
		return nil, err
	}
	return rs.DB, nil

}

func GetCache() (*cache.Client, error) {
	rs, err := Get()
	if err != nil {
		return nil, err
	}
	return rs.RC, nil
}