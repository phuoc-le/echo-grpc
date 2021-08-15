package config

import (
	"flag"
	"fmt"
	"grpc-echo/pkg/utils"
)

type Config struct {
	dbUser        string
	dbPswd        string
	dbHost        string
	dbPort        string
	dbName        string
	testDBHost    string
	testDBName    string
	apiPort       string
	migrate       string
	redisHost     string
	redisPort     string
	redisPassword string
	redisDB       int
}

func Get() *Config {
	conf := &Config{}

	flag.StringVar(&conf.dbUser, "dbuser", utils.GetStrEnv("POSTGRES_USER"), "DB user name")
	flag.StringVar(&conf.dbPswd, "dbpswd", utils.GetStrEnv("POSTGRES_PASSWORD"), "DB pass")
	flag.StringVar(&conf.dbPort, "dbport", utils.GetStrEnv("POSTGRES_PORT"), "DB port")
	flag.StringVar(&conf.dbHost, "dbhost", utils.GetStrEnv("POSTGRES_HOST"), "DB host")
	flag.StringVar(&conf.dbName, "dbname", utils.GetStrEnv("POSTGRES_DB"), "DB name")
	flag.StringVar(&conf.testDBHost, "testdbhost", utils.GetStrEnv("TEST_DB_HOST"), "test database host")
	flag.StringVar(&conf.testDBName, "testdbname", utils.GetStrEnv("TEST_DB_NAME"), "test database name")
	flag.StringVar(&conf.apiPort, "apiPort", utils.GetStrEnv("API_PORT"), "API Port")
	flag.StringVar(&conf.migrate, "migrate", "up", "specify if we should be migrating DB 'up' or 'down'")
	flag.StringVar(&conf.redisHost, "redisHost", utils.GetStrEnv("REDIS_HOST"), "Redis Host")
	flag.StringVar(&conf.redisPort, "redisPort", utils.GetStrEnv("REDIS_PORT"), "Redis Port")
	flag.StringVar(&conf.redisPassword, "redisPassword", utils.GetStrEnv("REDIS_PASSWORD"), "Redis Password")
	flag.IntVar(&conf.redisDB, "redisDB", utils.GetIntEnv("REDIS_DB"), "Redis Database")

	flag.Parse()

	return conf
}

func (c *Config) GetDBConnStr() string {
	return c.getDBConnStr(c.dbHost, c.dbName)
}

func (c *Config) GetTestDBConnStr() string {
	return c.getDBConnStr(c.testDBHost, c.testDBName)
}

func (c *Config) getDBConnStr(dbhost, dbname string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.dbUser,
		c.dbPswd,
		dbhost,
		c.dbPort,
		dbname,
	)
}

func (c *Config) GetAPIPort() string {
	return ":" + c.apiPort
}

func (c *Config) GetMigration() string {
	return c.migrate
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
