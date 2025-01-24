FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o jwtapi ./cmd/jwtapi/main.go

# Use a minimal Alpine image for the final stage
FROM alpine:latest

# Create the required directories
RUN mkdir -p /loki/rules

# Set permissions (optional)
RUN chmod -R 777 /loki/rules

WORKDIR /app

# Create the log directory and set permissions
RUN mkdir -p /var/log/jwtapi && chmod 777 /var/log/jwtapi

# Copy the binary from the builder stage
COPY --from=builder /app/jwtapi .

# Copy the log directory (if needed)
RUN mkdir -p /var/log/jwtapi

# Set the command to run the binary
CMD ["./jwtapi"]