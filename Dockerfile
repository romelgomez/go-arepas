# Stage 1: Build the Go application
FROM golang:1.23 AS builder

WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Download Prisma client dependencies
RUN go run github.com/steebchen/prisma-client-go generate

# Build the application
RUN go build -o main .

# Stage 2: Create the final runtime image
FROM debian:buster-slim

WORKDIR /app

# Copy the built application from the builder stage
COPY --from=builder /app/main .

# Expose the port your application listens on
EXPOSE 8080

# Start the application
CMD ["./main"]
