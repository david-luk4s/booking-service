package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectionDB() *gorm.DB {

	dsn := "host=localhost user=postgres password=postgres dbname=dbstudents port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Panic(err.Error())
	}
	DB.AutoMigrate()
	return DB
}
