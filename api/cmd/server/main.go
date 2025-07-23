package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"stock-api/internal/db"
	"stock-api/internal/handlers"
	"stock-api/internal/middleware"
)

func main() {
	// Load environment variables
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Connect to database
	if err := db.Connect(databaseURL); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Setup Gin
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// Add basic middleware
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(middleware.CORSConfig())

	// Setup routes
	api := r.Group("/api/v1")
	{
		// Stocks routes
		api.GET("/stocks", handlers.ListStocks)
		api.GET("/stocks/:ticker", handlers.GetStockDetails)

		// Recommendations routes
		api.GET("/recommendations", handlers.ListRecommendations)
		api.GET("/recommendations/:id", handlers.GetRecommendationByID)

		// Analytics routes
		api.GET("/analytics/summary", handlers.GetAnalyticsSummary)
		api.GET("/analytics/brokerages", handlers.GetTopBrokerages)

		// Health and migration routes
		api.GET("/health", handlers.HealthCheck)
		api.GET("/validate-migration", handlers.ValidateMigration)
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
