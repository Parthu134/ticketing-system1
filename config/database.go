package config

import (
	"log"
	"os"

	"ticketing-system/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	_ = godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set in environment")
	}
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	DB = db
	if err := db.AutoMigrate(&models.User{}, &models.Ticket{}, &models.Tag{}, &models.Subscription{}); err != nil {
		log.Fatal("failed to migrate tables", err)
	}
}
