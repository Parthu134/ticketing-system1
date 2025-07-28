package config

import (
	"log"

	"ticketing-system/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=Parthu732 dbname=ticketing port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	DB = db
	if err := db.AutoMigrate(&models.User{}, &models.Ticket{}, &models.Tag{}); err != nil {
		log.Fatal("failed to migrate tables", err)
	}
}
