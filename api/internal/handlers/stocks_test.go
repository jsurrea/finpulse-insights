package handlers_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"stock-api/internal/testutil"
)

func TestListStocksAndDetails(t *testing.T) {
	r := testutil.InitServer()

	// List endpoint
	w, body := testutil.Perform(r, "GET", "/api/v1/stocks?limit=10&page=1")
	testutil.AssertStatus(t, w.Code, 200)
	assert.Contains(t, body, `"AAPL"`)
	assert.Contains(t, body, `"MSFT"`)

	// Filter by brokerage
	_, body = testutil.Perform(r, "GET", "/api/v1/stocks?brokerage=Goldman+Sachs")
	assert.True(t, strings.Count(body, `"ticker"`) == 1)

	// Details OK
	w, body = testutil.Perform(r, "GET", "/api/v1/stocks/AAPL")
	testutil.AssertStatus(t, w.Code, 200)
	assert.Contains(t, body, `"latest_recommendation"`)

	// Details 404
	w, _ = testutil.Perform(r, "GET", "/api/v1/stocks/NONEXIST")
	testutil.AssertStatus(t, w.Code, 404)
}
