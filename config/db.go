package config

import (
	"log"
	"notification-hub/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, fallback to system env")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("✅ Database connected successfully")

	// Migrate models
	err = DB.AutoMigrate(
		&models.NotificationContext{},
		&models.EmailTemplate{},
		&models.Group{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
