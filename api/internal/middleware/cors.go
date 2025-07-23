package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORSConfig returns a Gin HandlerFunc for CORS config
func CORSConfig() gin.HandlerFunc {
	allowedOrigins := []string{"*"}
	// Optional: Allow CORS origins from env in production
	if envOrigins := os.Getenv("CORS_ALLOW_ORIGINS"); envOrigins != "" {
		// Comma-separated list
		allowedOrigins = []string{}
		for _, o := range splitAndTrim(envOrigins, ",") {
			if o != "" {
				allowedOrigins = append(allowedOrigins, o)
			}
		}
	}

	return cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

// Helper: split and trim env strings
func splitAndTrim(s, sep string) []string {
	parts := []string{}
	for _, p := range strings.Split(s, sep) {
		tp := strings.TrimSpace(p)
		if tp != "" {
			parts = append(parts, tp)
		}
	}
	return parts
}
