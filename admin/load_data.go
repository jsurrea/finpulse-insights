package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

// StockRecommendation represents a single stock recommendation item
type StockRecommendation struct {
    Ticker     string    `json:"ticker"`
    Company    string    `json:"company"`
    Brokerage  string    `json:"brokerage"`
    Action     string    `json:"action"`
    TargetFrom float64   `json:"target_from"`
    TargetTo   float64   `json:"target_to"`
    RatingFrom string    `json:"rating_from"`
    RatingTo   string    `json:"rating_to"`
    Time       time.Time `json:"time"`
}

// apiResponse represents the API response structure
type apiResponse struct {
    Items    []struct {
        Ticker     string `json:"ticker"`
        Company    string `json:"company"`
        Brokerage  string `json:"brokerage"`
        Action     string `json:"action"`
        TargetFrom string `json:"target_from"`
        TargetTo   string `json:"target_to"`
        RatingFrom string `json:"rating_from"`
        RatingTo   string `json:"rating_to"`
        Time       string `json:"time"`
    } `json:"items"`
    NextPage string `json:"next_page"`
}

const (
    dbURL      = "postgresql://root@localhost:26257/defaultdb?sslmode=disable"
    apiBaseURL = "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list"
)

// parseDollarValue converts a string like "$4.20" to float64
func parseDollarValue(s string) float64 {
    re := regexp.MustCompile(`[0-9.]+`)
    val := re.FindString(s)
    f, err := strconv.ParseFloat(val, 64)
    if err != nil {
        return 0
    }
    return f
}

// fetchPage retrieves a single page from the API
func fetchPage(token, nextPage string) (*apiResponse, error) {
    url := apiBaseURL
    if nextPage != "" {
        url += "?next_page=" + nextPage
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Authorization", "Bearer "+token)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{Timeout: 20 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
    }

    var apiResp apiResponse
    decoder := json.NewDecoder(resp.Body)
    if err := decoder.Decode(&apiResp); err != nil {
        return nil, err
    }
    return &apiResp, nil
}

// insertRecommendation inserts a recommendation into the database
func insertRecommendation(db *sql.DB, rec StockRecommendation) error {
    _, err := db.Exec(`
        INSERT INTO stock_recommendations
        (ticker, company, brokerage, action, target_from, target_to, rating_from, rating_to, time)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        ON CONFLICT DO NOTHING
    `,
        rec.Ticker, rec.Company, rec.Brokerage, rec.Action,
        rec.TargetFrom, rec.TargetTo, rec.RatingFrom, rec.RatingTo, rec.Time)
    return err
}

func main() {
    // Read API token from environment variable
    token := os.Getenv("API_TOKEN")
    if strings.TrimSpace(token) == "" {
        log.Fatal("API_TOKEN environment variable is required")
    }

    // Connect to CockroachDB
    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatalf("Failed to connect to DB: %v", err)
    }
    defer db.Close()

    _, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Scrape and insert all data
    nextPage := ""
    total := 0
    for {
        apiResp, err := fetchPage(token, nextPage)
        if err != nil {
            log.Fatalf("Failed to fetch data: %v", err)
        }
        for _, item := range apiResp.Items {
            t, err := time.Parse(time.RFC3339Nano, item.Time)
            if err != nil {
                log.Printf("Invalid time format: %v", item.Time)
                continue
            }
            rec := StockRecommendation{
                Ticker:     item.Ticker,
                Company:    item.Company,
                Brokerage:  item.Brokerage,
                Action:     item.Action,
                TargetFrom: parseDollarValue(item.TargetFrom),
                TargetTo:   parseDollarValue(item.TargetTo),
                RatingFrom: item.RatingFrom,
                RatingTo:   item.RatingTo,
                Time:       t,
            }
            if err := insertRecommendation(db, rec); err != nil {
                log.Printf("Failed to insert: %v", err)
            } else {
                total++
            }
        }
        if apiResp.NextPage == "" {
            break
        }
        nextPage = apiResp.NextPage
    }
    log.Printf("Done. Inserted %d recommendations.", total)
}
