package handlers_test

import (
	"testing"

	"stock-api/internal/testutil"
)

func TestHealth(t *testing.T) {
	r := testutil.InitServer()
	w, _ := testutil.Perform(r, "GET", "/api/v1/health")
	testutil.AssertStatus(t, w.Code, 200)
}
