package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"stock-api/internal/db"
	"stock-api/internal/models"
	"stock-api/internal/utils"
)

func ListStocks(c *gin.Context) {
	page := utils.ParsePage(c)
	limit := utils.ParseLimit(c)
	search := utils.SafeString(c.Query("search"))
	brokerage := utils.SafeString(c.Query("brokerage"))

	var stocks []struct {
		Ticker    string `json:"ticker"`
		Company   string `json:"company"`
		Brokerage string `json:"brokerage"`
	}

	query := db.GetDB().Model(&models.StockRecommendation{}).
		Select("DISTINCT ON (ticker) ticker, company, brokerage")

	if search != "" {
		searchPattern := "%" + strings.ToLower(search) + "%"
		query = query.Where("LOWER(ticker) LIKE ? OR LOWER(company) LIKE ?", searchPattern, searchPattern)
	}
	
	if brokerage != "" {
		query = query.Where("brokerage = ?", brokerage)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not count stocks"})
		return
	}

	err := query.
		Limit(limit).
		Offset((page - 1) * limit).
		Order("ticker ASC").
		Scan(&stocks).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch stocks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": stocks,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

func GetStockDetails(c *gin.Context) {
	ticker := utils.SafeString(strings.ToUpper(c.Param("ticker")))
	
	if ticker == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ticker is required"})
		return
	}

	var stock models.StockRecommendation
	err := db.GetDB().Where("ticker = ?", ticker).
		Order("time DESC").
		Limit(1).
		First(&stock).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "stock not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch stock details"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ticker":    stock.Ticker,
		"company":   stock.Company,
		"brokerage": stock.Brokerage,
		"latest_recommendation": gin.H{
			"action":      stock.Action,
			"rating_from": stock.RatingFrom,
			"rating_to":   stock.RatingTo,
			"target_from": stock.TargetFrom,
			"target_to":   stock.TargetTo,
			"time":        stock.Time.Format(time.RFC3339),
		},
		"history": []gin.H{
			{
				"id":             stock.ID.String(),
				"ticker":         stock.Ticker,
				"company":        stock.Company,
				"brokerage":      stock.Brokerage,
				"action":         stock.Action,
				"rating_from":    stock.RatingFrom,
				"rating_to":      stock.RatingTo,
				"target_from":    stock.TargetFrom,
				"target_to":      stock.TargetTo,
				"time":           stock.Time.Format(time.RFC3339),
				"recommendation": stock.Recommendation,
				"confidence":     stock.Confidence,
				"reason":         stock.Reason,
			},
		},
	})
}
