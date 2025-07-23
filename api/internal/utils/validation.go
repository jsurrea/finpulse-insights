package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const maxStringLength = 64

var allowedSortFields = map[string]bool{
	"time":         true,
	"ticker":       true,
	"company":      true,
	"brokerage":    true,
	"action":       true,
	"rating_to":    true,
	"confidence":   true,
}

var allowedSortDirections = map[string]bool{
	"asc":  true,
	"desc": true,
}

// ParsePage validates and returns a safe page number
func ParsePage(c *gin.Context) int {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		return 1
	}
	if page > 10000 { // Reasonable upper limit
		return 10000
	}
	return page
}

// ParseLimit validates and returns a safe limit
func ParseLimit(c *gin.Context) int {
	limitStr := c.DefaultQuery("limit", "20")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		return 20
	}
	if limit > 100 {
		return 100
	}
	return limit
}

// SafeString sanitizes and limits string input
func SafeString(input string) string {
	s := strings.TrimSpace(input)
	if len(s) > maxStringLength {
		s = s[:maxStringLength]
	}
	// Remove potentially dangerous characters for additional safety
	s = strings.ReplaceAll(s, "\x00", "")
	return s
}

// ParseSort validates sort parameters and returns safe field and direction
func ParseSort(sortParam string) (field, direction string) {
	field, direction = "time", "desc" // defaults
	
	if sortParam == "" {
		return
	}
	
	parts := strings.Split(sortParam, ":")
	if len(parts) != 2 {
		return
	}
	
	proposedField := strings.TrimSpace(parts[0])
	proposedDirection := strings.TrimSpace(strings.ToLower(parts[1]))
	
	if allowedSortFields[proposedField] && allowedSortDirections[proposedDirection] {
		field = proposedField
		direction = proposedDirection
	}
	
	return
}

// ValidateDate validates date string and returns parsed time
func ValidateDate(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, nil
	}
	
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format, expected YYYY-MM-DD")
	}
	
	// Additional validation: not too far in past/future
	now := time.Now()
	if date.After(now.AddDate(1, 0, 0)) {
		return time.Time{}, fmt.Errorf("date cannot be more than 1 year in the future")
	}
	if date.Before(now.AddDate(-10, 0, 0)) {
		return time.Time{}, fmt.Errorf("date cannot be more than 10 years in the past")
	}
	
	return date, nil
}

// Percent calculates percentage safely
func Percent(part, total int64) float64 {
	if total == 0 {
		return 0.0
	}
	return float64(part) / float64(total) * 100.0
}
