package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Listening on 0.0.0.0:%s", port)
    r := gin.Default()

    initDB()

    // CORS
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))
    
    // Routes
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "healthy"})
    })
    r.GET("/stocks", listStocks)
    r.GET("/stocks/:ticker", getStockDetails)
    r.GET("/recommendations", listRecommendations)
    r.GET("/recommendations/:id", getRecommendationByID)
    r.GET("/analytics/summary", getAnalyticsSummary)
    r.GET("/analytics/brokerages", getTopBrokerages)
    r.GET("/health", healthCheck)

    if err := r.Run("0.0.0.0:" + port); err != nil {
        log.Fatalf("Failed to run: %v", err)
    }
}
