# -------- Build stage --------
FROM golang:1.25-alpine AS builder

# Install git (needed for fetching dependencies)
RUN apk add --no-cache git

WORKDIR /app

# Copy mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 go build -o app ./cmd

# -------- Final stage --------
FROM alpine:latest

# Install CA certificates for TLS support (important for gRPC!)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/app .
RUN chmod +x ./app

# Expose gRPC port
EXPOSE 50051

# Run the binary
CMD ["./app"]
