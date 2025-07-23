package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"stock-api/internal/models"
)

var DB *gorm.DB

func Connect(databaseURL string) error {
	var err error
	
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	
	DB, err = gorm.Open(postgres.Open(databaseURL), config)
	if err != nil {
		return err
	}
	
	// Configure connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	
	// Auto migrate
	err = DB.AutoMigrate(&models.StockRecommendation{})
	if err != nil {
		return err
	}
	
	log.Println("Database connected and migrated successfully")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
