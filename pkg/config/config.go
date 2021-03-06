package config

import (
	"flag"
	"grpc-echo/pkg/utils"
)

type Config struct {
	dbUrl		  string
	apiPort       string
	grpcPort      string
	redisHost     string
	redisPort     string
	redisPassword string
	redisDB       int
}

func Get() *Config {
	conf := &Config{}

	flag.StringVar(&conf.dbUrl, "dbUrl", utils.GetStrEnv("POSTGRES_URL"), "DB Url")
	flag.StringVar(&conf.apiPort, "apiPort", utils.GetStrEnv("API_PORT"), "API Port")
	flag.StringVar(&conf.grpcPort, "grpcPort", utils.GetStrEnv("API_PORT"), "API Port")
	flag.StringVar(&conf.redisHost, "redisHost", utils.GetStrEnv("REDIS_HOST"), "Redis Host")
	flag.StringVar(&conf.redisPort, "redisPort", utils.GetStrEnv("REDIS_PORT"), "Redis Port")
	flag.StringVar(&conf.redisPassword, "redisPassword", utils.GetStrEnv("REDIS_PASSWORD"), "Redis Password")
	flag.IntVar(&conf.redisDB, "redisDB", utils.GetIntEnv("REDIS_DB"), "Redis Database")
	flag.Parse()

	return conf
}

func (c *Config) GetDBConnStr() string {
	return c.dbUrl
}

func (c *Config) GetAPIPort() string {
	return ":" + c.apiPort
}

func (c *Config) GetGrpcPort() string {
	return ":" + c.grpcPort
}


func (c *Config) GetRedisUrl() string {
	return c.redisHost + ":" + c.redisPort
}

func (c *Config) GetRedisDB() int {
	return c.redisDB
}
func (c *Config) GetRedisPassword() string {
	return c.redisPassword
}