package database

import (
	"gim-api-rest-go/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func ConnectDatabase() {
	stringConnection := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	
	DB, err = gorm.Open(postgres.Open(stringConnection))

	if err != nil {
		log.Panic("Error to connect with database")
	}

	DB.AutoMigrate(&models.Student{})
}