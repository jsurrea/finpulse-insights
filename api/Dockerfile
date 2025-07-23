# ---------- Build Stage ----------
FROM golang:alpine AS builder

WORKDIR /app

# Copia go.mod y go.sum y descarga dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto del proyecto y compila binario estático
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o stocks-api ./cmd/server

# ---------- Final Image ----------
FROM gcr.io/distroless/static

WORKDIR /

COPY --from=builder /app/stocks-api /stocks-api

EXPOSE 8080

ENV PORT=8080

ENTRYPOINT ["/stocks-api"]
