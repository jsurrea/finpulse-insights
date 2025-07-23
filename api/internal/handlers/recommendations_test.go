package handlers_test

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"stock-api/internal/db"
	"stock-api/internal/models"
	"stock-api/internal/testutil"
)

func TestRecommendationsEndpoints(t *testing.T) {
	r := testutil.InitServer()

	// List all
	w, body := testutil.Perform(r, "GET", "/api/v1/recommendations")
	testutil.AssertStatus(t, w.Code, 200)
	assert.Contains(t, body, `"AAPL"`)

	// Filter
	_, body = testutil.Perform(r, "GET", "/api/v1/recommendations?ticker=MSFT")
	assert.True(t, strings.Count(body, `"MSFT"`) == 1)

	// Get by ID valid
	var rec models.StockRecommendation
	db.GetDB().First(&rec, "ticker = ?", "AAPL")
	w, _ = testutil.Perform(r, "GET", "/api/v1/recommendations/"+rec.ID.String())
	testutil.AssertStatus(t, w.Code, 200)

	// Invalid UUID
	w, _ = testutil.Perform(r, "GET", "/api/v1/recommendations/not-a-uuid")
	testutil.AssertStatus(t, w.Code, 400)

	// Not found
	w, _ = testutil.Perform(r, "GET", "/api/v1/recommendations/"+uuid.NewString())
	testutil.AssertStatus(t, w.Code, 404)
}
