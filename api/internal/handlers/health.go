package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"stock-api/internal/db"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

func HealthCheck(c *gin.Context) {
	database := db.GetDB()
	if database == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   "error", 
			"database": "not connected",
		})
		return
	}

	sqlDB, err := database.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   "error", 
			"database": "not connected",
		})
		return
	}

	err = sqlDB.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   "error", 
			"database": "not connected",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"database": "connected",
		"uptime":   time.Since(startTime).Seconds(),
	})
}
