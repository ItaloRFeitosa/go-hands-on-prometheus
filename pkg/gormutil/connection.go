package gormutil

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DefaultDatabaseURL = os.Getenv("DEFAULT_DATABASE_URL")
	defaultDB          *gorm.DB
)

func Connect() *gorm.DB {
	if defaultDB != nil {
		return defaultDB
	}

	newDB, err := gorm.Open(postgres.Open(DefaultDatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	defaultDB = newDB

	return defaultDB
}
