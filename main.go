package main

import (
	"log"
	"notification-hub/config"
	"notification-hub/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}
	// Initialize database
	config.InitDB()

	r := gin.Default()

	// Register all routes
	routes.RegisterRoutes(r)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
