package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
    initDB()

    r := gin.Default()

    // Routes
    r.GET("/stocks", listStocks)
    r.GET("/stocks/:ticker", getStockDetails)
    r.GET("/recommendations", listRecommendations)
    r.GET("/recommendations/:id", getRecommendationByID)
    r.GET("/analytics/summary", getAnalyticsSummary)
    r.GET("/health", healthCheck)

    r.Run(":8081")
}
