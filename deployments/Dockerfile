# Build stage
FROM golang:1.20 AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd/myapp

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/myapp .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./myapp"]