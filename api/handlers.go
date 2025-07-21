package main

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func listStocks(c *gin.Context) {
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
    if limit > 100 {
        limit = 100
    }
    search := c.Query("search")
    brokerage := c.Query("brokerage")

    var stocks []struct {
        Ticker    string `json:"ticker"`
        Company   string `json:"company"`
        Brokerage string `json:"brokerage"`
    }

    query := db.Model(&StockRecommendation{}).
        Select("DISTINCT ON (ticker) ticker, company, brokerage")

    if search != "" {
        likeSearch := "%" + strings.ToLower(search) + "%"
        query = query.Where("LOWER(ticker) LIKE ? OR LOWER(company) LIKE ?", likeSearch, likeSearch)
    }
    if brokerage != "" {
        query = query.Where("brokerage = ?", brokerage)
    }

    var total int64
    query.Count(&total)

    err := query.
        Limit(limit).
        Offset((page - 1) * limit).
        Order("ticker ASC").
        Scan(&stocks).Error

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

func getStockDetails(c *gin.Context) {
    ticker := c.Param("ticker")

    var stock StockRecommendation
    err := db.Where("ticker = ?", ticker).
        Order("time DESC").
        Limit(1).
        First(&stock).Error

    if err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "stock not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
            "time":        stock.Time,
        },
    })
}

func listRecommendations(c *gin.Context) {
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
    if limit > 100 {
        limit = 100
    }

    ticker := c.Query("ticker")
    company := c.Query("company")
    brokerage := c.Query("brokerage")
    action := c.Query("action")
    rating := c.Query("rating")
    dateFrom := c.Query("date_from")
    dateTo := c.Query("date_to")
    sort := c.DefaultQuery("sort", "time:desc")

    query := db.Model(&StockRecommendation{})

    if ticker != "" {
        query = query.Where("ticker = ?", ticker)
    }
    if company != "" {
        query = query.Where("company = ?", company)
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
    if dateFrom != "" {
        if t, err := time.Parse("2006-01-02", dateFrom); err == nil {
            query = query.Where("time >= ?", t)
        }
    }
    if dateTo != "" {
        if t, err := time.Parse("2006-01-02", dateTo); err == nil {
            query = query.Where("time <= ?", t)
        }
    }

    var total int64
    query.Count(&total)

    // Sorting
    sortField := "time"
    sortOrder := "desc"
    if parts := strings.Split(sort, ":"); len(parts) == 2 {
        sortField = parts[0]
        sortOrder = parts[1]
    }
    orderStr := sortField + " " + sortOrder

    var recs []StockRecommendation
    err := query.Order(orderStr).
        Limit(limit).
        Offset((page - 1) * limit).
        Find(&recs).Error

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

func getRecommendationByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := uuid.Parse(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
        return
    }

    var rec StockRecommendation
    err = db.First(&rec, "id = ?", id).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "recommendation not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, rec)
}

func getAnalyticsSummary(c *gin.Context) {
    var totalStocks int64
    var totalRecommendations int64
    var buyCount int64
    var sellCount int64
    var holdCount int64
    var avgConfidence float64

    db.Model(&StockRecommendation{}).Distinct("ticker").Count(&totalStocks)
    db.Model(&StockRecommendation{}).Count(&totalRecommendations)
    db.Model(&StockRecommendation{}).Where("recommendation = ?", "BUY").Count(&buyCount)
    db.Model(&StockRecommendation{}).Where("recommendation = ?", "SELL").Count(&sellCount)
    db.Model(&StockRecommendation{}).Where("recommendation = ?", "HOLD").Count(&holdCount)
    db.Model(&StockRecommendation{}).Select("AVG(confidence)").Scan(&avgConfidence)

    c.JSON(http.StatusOK, gin.H{
        "total_stocks":         totalStocks,
        "total_recommendations": totalRecommendations,
        "buy_percentage":       percent(buyCount, totalRecommendations),
        "sell_percentage":      percent(sellCount, totalRecommendations),
        "hold_percentage":      percent(holdCount, totalRecommendations),
        "average_confidence":   avgConfidence,
    })
}

func getTopBrokerages(c *gin.Context) {
    type Brokerage struct {
        Name       string `json:"name"`
        StockCount int    `json:"stockCount"`
    }
    var results []Brokerage

    err := db.Model(&StockRecommendation{}).
        Select("brokerage AS name, COUNT(DISTINCT ticker) AS stock_count").
        Group("brokerage").
        Order("stock_count DESC").
        Scan(&results).Error

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, results)
}

func healthCheck(c *gin.Context) {
    sqlDB, err := db.DB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "database": "not connected"})
        return
    }
    err = sqlDB.Ping()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "database": "not connected"})
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "status":   "ok",
        "database": "connected",
        "uptime":   time.Now().Unix(),
    })
}

func validateMigration(c *gin.Context) {
    var stats struct {
        Total           int64   `json:"total"`
        BuyCount        int64   `json:"buy_count"`
        SellCount       int64   `json:"sell_count"`
        HoldCount       int64   `json:"hold_count"`
        AvgConfidence   float64 `json:"avg_confidence"`
        EmptyRecs       int64   `json:"empty_recommendations"`
        ZeroConfidence  int64   `json:"zero_confidence"`
    }
    
    db.Model(&StockRecommendation{}).Count(&stats.Total)
    db.Model(&StockRecommendation{}).Where("recommendation = 'BUY'").Count(&stats.BuyCount)
    db.Model(&StockRecommendation{}).Where("recommendation = 'SELL'").Count(&stats.SellCount)
    db.Model(&StockRecommendation{}).Where("recommendation = 'HOLD'").Count(&stats.HoldCount)
    db.Model(&StockRecommendation{}).Where("recommendation = '' OR recommendation IS NULL").Count(&stats.EmptyRecs)
    db.Model(&StockRecommendation{}).Where("confidence = 0").Count(&stats.ZeroConfidence)
    db.Model(&StockRecommendation{}).Select("AVG(confidence)").Scan(&stats.AvgConfidence)
    
    c.JSON(http.StatusOK, gin.H{
        "migration_stats": stats,
        "status": "validation_complete",
    })
}

