package main

import (
	"os"

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
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database: " + err.Error())
    }

    // Enable UUID extension for CockroachDB
    db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

    // Auto migrate the schema
    err = db.AutoMigrate(&StockRecommendation{})
    if err != nil {
        panic("failed to migrate database: " + err.Error())
    }
}