package main

import (
	"log"
	"notification-hub/config"
	"notification-hub/handlers"
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

	// Initialize handlers
	contextHandler := handlers.NewContextHandler(config.DB)

	// Define routes
	r.POST("/contexts", contextHandler.Create)
	r.GET("/contexts", contextHandler.List)
	r.GET("/contexts/:id", contextHandler.Get)
	r.PUT("/contexts/:id", contextHandler.Update)
	r.DELETE("/contexts/:id", contextHandler.Delete)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
