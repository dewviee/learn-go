package database

import (
	"log"

	"github.com/dewviee/learn-go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(conf config.DBConfig) *gorm.DB {
	postgresConf, gormConf := config.GetDBConfig(conf)

	db, err := gorm.Open(postgres.New(postgresConf), &gormConf)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return db
}
