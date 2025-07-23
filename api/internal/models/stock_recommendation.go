package models

import (
	"time"

	"github.com/google/uuid"
)

type StockRecommendation struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Ticker         string    `gorm:"not null;index" json:"ticker"`
	Company        string    `gorm:"not null" json:"company"`
	Brokerage      string    `gorm:"not null;index" json:"brokerage"`
	Action         string    `gorm:"not null" json:"action"`
	TargetFrom     *float64  `json:"target_from"`
	TargetTo       *float64  `json:"target_to"`
	RatingFrom     *string   `json:"rating_from"`
	RatingTo       string    `gorm:"not null" json:"rating_to"`
	Time           time.Time `gorm:"not null;index" json:"time"`
	
	// Calculated fields
	Recommendation string  `gorm:"not null;default:'HOLD'" json:"recommendation"`
	Confidence     float64 `gorm:"not null;default:0.5" json:"confidence"`
	Score          int     `gorm:"default:0" json:"score"`
	Reason         string  `gorm:"type:text" json:"reason"`
	
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (StockRecommendation) TableName() string {
	return "stock_recommendations"
}
