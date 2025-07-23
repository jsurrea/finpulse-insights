package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectInvalid(t *testing.T) {
	// Use a very unlikely invalid DSN
	err := Connect("postgres://bad:bad@bad/bad?sslmode=disable")
	assert.Error(t, err)
}

func TestConnectMissingEnv(t *testing.T) {
	// Clear DATABASE_URL env var
	os.Unsetenv("DATABASE_URL")

	err := Connect("")
	assert.Error(t, err)
}
