package main

import (
	"time"

	"github.com/google/uuid"
)

type StockRecommendation struct {
    ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
    Ticker      string    `gorm:"not null" json:"ticker"`
    Company     string    `gorm:"not null" json:"company"`
    Brokerage   string    `gorm:"not null" json:"brokerage"`
    Action      string    `gorm:"not null" json:"action"`
    TargetFrom  float64   `json:"target_from"`
    TargetTo    float64   `json:"target_to"`
    RatingFrom  string    `json:"rating_from"`
    RatingTo    string    `json:"rating_to"`
    Time        time.Time `gorm:"not null" json:"time"`
    Score       int       `json:"score"`          // Optional: for recommendation algorithm
    Recommendation string  `json:"recommendation"` // BUY/SELL/HOLD
    Confidence  float64   `json:"confidence"`     // 0.0 to 1.0
    Reason      string    `json:"reason"`         // Explanation text
}
