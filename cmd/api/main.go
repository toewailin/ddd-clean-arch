package main

import (
	"ddd/config"
	"ddd/internal/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()      // Load config
	config.ConnectDB(cfg)     // Get the database connection

	// Create a new Gin router
	r := gin.Default()

	// Register routes from the router package
	router.RegisterRoutes(r, config.DB)    // Pass the *gorm.DB connection here

	// Start the server
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
