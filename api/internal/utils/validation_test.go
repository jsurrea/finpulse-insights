package utils

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Helper para crear un Gin Context con query params
func createTestCtx(queryParams map[string]string) *gin.Context {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	values := url.Values{}
	for k, v := range queryParams {
		values.Set(k, v)
	}

	c.Request = &http.Request{
		URL: &url.URL{RawQuery: values.Encode()},
	}
	return c
}

func TestParsePage(t *testing.T) {
	cases := []struct {
		query    string
		expected int
	}{
		{"5", 5},
		{"0", 1},
		{"-3", 1},
		{"15000", 10000},
		{"nonint", 1},
		{"", 1},
	}

	for _, c := range cases {
		ctx := createTestCtx(map[string]string{"page": c.query})
		val := ParsePage(ctx)
		assert.Equal(t, c.expected, val, "query: %s", c.query)
	}
}

func TestParseLimit(t *testing.T) {
	cases := []struct {
		query    string
		expected int
	}{
		{"25", 25}, {"0", 20}, {"-3", 20}, {"150", 100}, {"nonint", 20}, {"", 20},
	}

	for _, c := range cases {
		ctx := createTestCtx(map[string]string{"limit": c.query})
		val := ParseLimit(ctx)
		assert.Equal(t, c.expected, val, "query: %s", c.query)
	}
}

func TestSafeString(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{" hello ", "hello"},
		{"text\x00more", "textmore"},
		{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+{}|:\"<>?", "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"},
		{"normal", "normal"},
	}

	for _, tt := range tests {
		got := SafeString(tt.input)
		if len(tt.input) > maxStringLength { // maxStringLength=64
			assert.Len(t, got, maxStringLength)
		} else {
			assert.Equal(t, tt.want, got)
		}
	}
}


func TestParseSort(t *testing.T) {
	tests := []struct {
		input string
		wantF string
		wantD string
	}{
		{"ticker:asc", "ticker", "asc"},
		{"time:desc", "time", "desc"},
		{"ticker:bad", "time", "desc"},
		{"badfield:asc", "time", "desc"},
		{"", "time", "desc"},
	}

	for _, tt := range tests {
		f, d := ParseSort(tt.input)
		assert.Equal(t, tt.wantF, f)
		assert.Equal(t, tt.wantD, d)
	}
}

func TestValidateDate(t *testing.T) {
	tests := []struct {
		input      string
		shouldFail bool
	}{
		{"2025-01-02", false},
		{"", false},
		{"01-02-2025", true},
		{"bad", true},
		{"2099-01-01", true},
		{"1900-01-01", true},
	}

	for _, tt := range tests {
		_, err := ValidateDate(tt.input)
		if tt.shouldFail {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestPercent(t *testing.T) {
	assert.InDelta(t, 0.0, Percent(0, 0), 1e-9)
	assert.InDelta(t, 50.0, Percent(50, 100), 1e-9)
	assert.InDelta(t, 33.33333333333333, Percent(1, 3), 1e-9) // tolerancia 1e-9
}

