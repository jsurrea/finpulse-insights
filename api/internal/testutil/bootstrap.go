package testutil

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"stock-api/internal/db"
	"stock-api/internal/handlers"
	"stock-api/internal/middleware"
	"stock-api/internal/models"
)

// InitServer connects to the DATABASE_URL already exported by scripts/test.sh
// It returns a fully wired Gin engine and clears the table before each test suite.
func InitServer() *gin.Engine {
	// Connect once per package
	if db.GetDB() == nil {
		if err := db.Connect(os.Getenv("DATABASE_URL")); err != nil {
			panic(err)
		}
	}

	// Clean table
	db.GetDB().Exec("TRUNCATE TABLE stock_recommendations")

	// Seed two rows covering BUY, SELL, HOLD
	now := time.Now()
	db.GetDB().Create([]models.StockRecommendation{
		{
			ID:     uuid.New(), Ticker: "AAPL", Company: "Apple Inc.",
			Brokerage: "Goldman Sachs", Action: "upgraded by",
			RatingTo: "Buy", Time: now, Recommendation: "BUY", Confidence: 0.9,
		},
		{
			ID:     uuid.New(), Ticker: "MSFT", Company: "Microsoft",
			Brokerage: "Morgan Stanley", Action: "downgraded by",
			RatingTo: "Sell", Time: now, Recommendation: "SELL", Confidence: 0.6,
		},
	})

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(gin.Recovery(), middleware.CORSConfig())

	api := r.Group("/api/v1")
	{
		api.GET("/stocks", handlers.ListStocks)
		api.GET("/stocks/:ticker", handlers.GetStockDetails)

		api.GET("/recommendations", handlers.ListRecommendations)
		api.GET("/recommendations/:id", handlers.GetRecommendationByID)

		api.GET("/analytics/summary", handlers.GetAnalyticsSummary)
		api.GET("/analytics/brokerages", handlers.GetTopBrokerages)

		api.GET("/health", handlers.HealthCheck)
		api.GET("/validate-migration", handlers.ValidateMigration)
	}

	return r
}

// Helper to perform a request and return recorder + body string
func Perform(r *gin.Engine, method, path string) (*httptest.ResponseRecorder, string) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, nil)
	r.ServeHTTP(w, req)
	return w, w.Body.String()
}

func AssertStatus(t TestingT, got, want int) {
	if got != want {
		log.Fatalf("unexpected status: got %d, want %d", got, want)
	}
}

type TestingT interface {
	Helper()
}
