package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Name     string `env:"DB_NAME "`
	User     string `env:"DB_USER "`
	Password string `env:"DB_PASSWORD "`
	Host     string `env:"DB_HOST "`
	Port     string `env:"DB_PORT "`
	Params   string `env:"DB_PARAMS "`
}

func GetDBConfig(conf DBConfig) (postgres.Config, gorm.Config) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s %s", conf.Host, conf.Port, conf.User, conf.Password, conf.Name, conf.Params)

	return postgres.Config{
		DSN: dsn,
	}, gorm.Config{}
}
