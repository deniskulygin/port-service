FROM golang:1.19-alpine AS builder
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/port-service/main.go
FROM alpine:latest

RUN apk upgrade --no-cache

# Reduce image size
RUN rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

# Avoid running code as a root user
RUN adduser -D appuser
USER appuser

WORKDIR /app

# Copy only the necessary files from the builder stage
COPY --from=builder /app/app .

# Set any environment variables required by the application
ENV HTTP_ADDR=:8080

# Expose the port that the application listens on
EXPOSE 8080

# Run the binary when the container starts
CMD ["./app"]
