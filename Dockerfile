FROM golang:1.24-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make

# Set working directory
WORKDIR /app

# Copy source code
COPY . .

# Build using the original Makefile
RUN make

# Force new deployment

# Final stage
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache ca-certificates

# Copy binary from builder
COPY --from=builder /app/dist/proxy /usr/local/bin/modlishka

# Copy config
COPY config.json /app/config.json

# Set working directory
WORKDIR /app

# Expose port
EXPOSE 8080

# Run the application
CMD ["modlishka", "-config", "config.json"]