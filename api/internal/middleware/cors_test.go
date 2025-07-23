package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCORSDefault(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(CORSConfig())
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Origin", "http://example.com")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
}

func TestCORSWithEnv(t *testing.T) {
	gin.SetMode(gin.TestMode)
	os.Setenv("CORS_ALLOW_ORIGINS", "http://example.com,https://foo.com")
	defer os.Unsetenv("CORS_ALLOW_ORIGINS")

	r := gin.New()
	r.Use(CORSConfig())
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Origin", "http://example.com")

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "http://example.com", w.Header().Get("Access-Control-Allow-Origin"))
}

func TestSplitAndTrim(t *testing.T) {
	input := "foo, bar ,baz,,"
	expected := []string{"foo", "bar", "baz"}

	result := splitAndTrim(input, ",")

	assert.Equal(t, expected, result)
}
