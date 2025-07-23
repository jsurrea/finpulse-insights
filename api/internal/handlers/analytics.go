package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"stock-api/internal/db"
	"stock-api/internal/models"
	"stock-api/internal/utils"
)

func GetAnalyticsSummary(c *gin.Context) {
	var totalStocks int64
	var totalRecommendations int64
	var buyCount int64
	var sellCount int64
	var holdCount int64
	var avgConfidence float64

	database := db.GetDB()

	// Execute queries with error handling
	if err := database.Model(&models.StockRecommendation{}).Distinct("ticker").Count(&totalStocks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not count total stocks"})
		return
	}

	if err := database.Model(&models.StockRecommendation{}).Count(&totalRecommendations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not count total recommendations"})
		return
	}

	if err := database.Model(&models.StockRecommendation{}).Where("recommendation = ?", "BUY").Count(&buyCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not count BUY recommendations"})
		return
	}

	if err := database.Model(&models.StockRecommendation{}).Where("recommendation = ?", "SELL").Count(&sellCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not count SELL recommendations"})
		return
	}

	if err := database.Model(&models.StockRecommendation{}).Where("recommendation = ?", "HOLD").Count(&holdCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not count HOLD recommendations"})
		return
	}

	if err := database.Model(&models.StockRecommendation{}).Select("AVG(confidence)").Scan(&avgConfidence).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not calculate average confidence"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_stocks":           totalStocks,
		"total_recommendations":  totalRecommendations,
		"buy_percentage":         utils.Percent(buyCount, totalRecommendations),
		"sell_percentage":        utils.Percent(sellCount, totalRecommendations),
		"hold_percentage":        utils.Percent(holdCount, totalRecommendations),
		"average_confidence":     avgConfidence,
	})
}

func GetTopBrokerages(c *gin.Context) {
	type Brokerage struct {
		Name       string `json:"name"`
		StockCount int    `json:"stockCount"`
	}
	var results []Brokerage

	err := db.GetDB().Model(&models.StockRecommendation{}).
		Select("brokerage AS name, COUNT(DISTINCT ticker) AS stock_count").
		Group("brokerage").
		Order("stock_count DESC").
		Scan(&results).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch top brokerages"})
		return
	}

	c.JSON(http.StatusOK, results)
}

func ValidateMigration(c *gin.Context) {
	var stats struct {
		Total          int64   `json:"total"`
		BuyCount       int64   `json:"buy_count"`
		SellCount      int64   `json:"sell_count"`
		HoldCount      int64   `json:"hold_count"`
		AvgConfidence  float64 `json:"avg_confidence"`
		EmptyRecs      int64   `json:"empty_recommendations"`
		ZeroConfidence int64   `json:"zero_confidence"`
	}

	database := db.GetDB()

	database.Model(&models.StockRecommendation{}).Count(&stats.Total)
	database.Model(&models.StockRecommendation{}).Where("recommendation = 'BUY'").Count(&stats.BuyCount)
	database.Model(&models.StockRecommendation{}).Where("recommendation = 'SELL'").Count(&stats.SellCount)
	database.Model(&models.StockRecommendation{}).Where("recommendation = 'HOLD'").Count(&stats.HoldCount)
	database.Model(&models.StockRecommendation{}).Where("recommendation = '' OR recommendation IS NULL").Count(&stats.EmptyRecs)
	database.Model(&models.StockRecommendation{}).Where("confidence = 0").Count(&stats.ZeroConfidence)
	database.Model(&models.StockRecommendation{}).Select("AVG(confidence)").Scan(&stats.AvgConfidence)

	c.JSON(http.StatusOK, gin.H{
		"migration_stats": stats,
		"status":          "validation_complete",
	})
}
