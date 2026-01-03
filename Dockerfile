# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies for CGO (required by go-sqlite3)
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=1 GOOS=linux go build -o server .

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/server .
COPY --from=builder /app/graphiql.html .

EXPOSE 8080

CMD ["./server"]
