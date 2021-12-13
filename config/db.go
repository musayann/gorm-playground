package config

import (
	models "musayann/gorm-playground/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func ConnectDB() {
	dsn := "host=127.0.0.1 user=postgres dbname=gorm port=5432 sslmode=disable"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	Db.AutoMigrate(&models.Product{})
	Db.AutoMigrate(&models.Sale{})
}
