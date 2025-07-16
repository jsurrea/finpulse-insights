package main

import (
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgresql://root@localhost:26257/defaultdb?sslmode=disable"
	}

	var err error
	maxAttempts := 10

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		log.Printf("Connecting to DB (attempt %d/%d)...", attempt, maxAttempts)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Printf("Failed to open DB: %v", err)
		} else {
			// Confirm connection via ping
			sqlDB, pingErr := db.DB()
			if pingErr == nil {
				pingErr = sqlDB.Ping()
			}
			if pingErr == nil {
				log.Println("Successfully connected to database!")
				break
			} else {
				log.Printf("Ping failed: %v", pingErr)
			}
		}

		if attempt == maxAttempts {
			log.Fatalf("Could not connect to database after %d attempts", maxAttempts)
		}

		time.Sleep(5 * time.Second)
	}

	// Run extension enabling (harmless if it fails, CockroachDB may ignore this)
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	// Auto migrate
	if err := db.AutoMigrate(&StockRecommendation{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
