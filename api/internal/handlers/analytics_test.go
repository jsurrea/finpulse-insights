package handlers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"stock-api/internal/testutil"
)

func TestAnalyticsEndpoints(t *testing.T) {
	r := testutil.InitServer()

	// Summary
	w, body := testutil.Perform(r, "GET", "/api/v1/analytics/summary")
	testutil.AssertStatus(t, w.Code, 200)
	assert.Contains(t, body, `"total_stocks":2`)
	assert.Contains(t, body, `"average_confidence"`)

	// Brokerages
	w, body = testutil.Perform(r, "GET", "/api/v1/analytics/brokerages")
	testutil.AssertStatus(t, w.Code, 200)
	assert.Contains(t, body, "Goldman Sachs")
}
