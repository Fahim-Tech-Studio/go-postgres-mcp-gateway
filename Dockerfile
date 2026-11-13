
# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum* ./
RUN go mod download

# Copy source code
COPY . .

# Build the server binary with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -o mcp-gateway ./cmd/server/main.go

# Stage 2: Runtime Production Image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy compiled binary from builder stage
COPY --from=builder /app/mcp-gateway .

# Expose port for SSE/HTTP transport (if applicable)
EXPOSE 8080

# Run the gateway binary
CMD ["./mcp-gateway"]
