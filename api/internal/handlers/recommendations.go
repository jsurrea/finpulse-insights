package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"stock-api/internal/db"
	"stock-api/internal/models"
	"stock-api/internal/utils"
)

func ListRecommendations(c *gin.Context) {
	page := utils.ParsePage(c)
	limit := utils.ParseLimit(c)

	ticker := utils.SafeString(c.Query("ticker"))
	company := utils.SafeString(c.Query("company"))
	brokerage := utils.SafeString(c.Query("brokerage"))
	action := utils.SafeString(c.Query("action"))
	rating := utils.SafeString(c.Query("rating"))
	dateFromStr := c.Query("date_from")
	dateToStr := c.Query("date_to")
	sortParam := c.DefaultQuery("sort", "time:desc")

	// Validate dates
	dateFrom, err := utils.ValidateDate(dateFromStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid date_from: %v", err)})
		return
	}

	dateTo, err := utils.ValidateDate(dateToStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid date_to: %v", err)})
		return
	}

	// Validate and parse sort
	sortField, sortDirection := utils.ParseSort(sortParam)
	orderStr := fmt.Sprintf("%s %s", sortField, sortDirection)

	query := db.GetDB().Model(&models.StockRecommendation{})

	if ticker != "" {
		tickerPattern := "%" + strings.ToLower(ticker) + "%"
		query = query.Where("LOWER(ticker) LIKE ?", tickerPattern)
	}
	if company != "" {
		companyPattern := "%" + strings.ToLower(company) + "%"
		query = query.Where("LOWER(company) LIKE ?", companyPattern)
	}
	if brokerage != "" {
		query = query.Where("brokerage = ?", brokerage)
	}
	if action != "" {
		query = query.Where("action = ?", action)
	}
	if rating != "" {
		query = query.Where("rating_to = ?", rating)
	}
	if !dateFrom.IsZero() {
		query = query.Where("time >= ?", dateFrom)
	}
	if !dateTo.IsZero() {
		query = query.Where("time <= ?", dateTo)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not count recommendations"})
		return
	}

	var recs []models.StockRecommendation
	err = query.
		Order(orderStr).
		Limit(limit).
		Offset((page - 1) * limit).
		Find(&recs).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch recommendations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": recs,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

func GetRecommendationByID(c *gin.Context) {
	idParam := c.Param("id")
	
	// Validate UUID length first
	if len(idParam) > 36 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID format"})
		return
	}
	
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID format"})
		return
	}

	var rec models.StockRecommendation
	err = db.GetDB().First(&rec, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "recommendation not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch recommendation"})
		}
		return
	}

	c.JSON(http.StatusOK, rec)
}
