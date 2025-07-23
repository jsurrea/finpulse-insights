#!/usr/bin/env bash
set -euo pipefail

# ---- 1. Launch disposable Postgres ----
export TEST_PG_PORT=55432
docker run --rm -d \
  --name stock-api-test-db \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=stocks \
  -p ${TEST_PG_PORT}:5432 \
  postgres:alpine > /dev/null

# ---- 2. Wait until Postgres is ready ----
echo -n "Waiting for Postgres "
until docker exec stock-api-test-db pg_isready -U postgres -d stocks > /dev/null 2>&1; do
  echo -n "."
  sleep 0.5
done
echo " ready"

# ---- 3. Run tests with full coverage ----
export DATABASE_URL="postgres://postgres:postgres@localhost:${TEST_PG_PORT}/stocks?sslmode=disable"
export GIN_MODE=release # silence Gin debug logs
go test -coverprofile=coverage.out ./...

# ---- 4. Show per-file coverage summary ----
go tool cover -func=coverage.out

# ---- 5. Stop container ----
docker stop stock-api-test-db > /dev/null
