# Stage 1: Build the Go application
FROM golang:1.23 AS builder

WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Generate Prisma client
RUN go run github.com/steebchen/prisma-client-go generate

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Create the final runtime image
FROM alpine:latest AS runtime

WORKDIR /app

# Copy the built application from the builder stage
COPY --from=builder /app/main .

# Copy the .env file
COPY --from=builder /app/.env.example .env

# Install any other runtime dependencies if needed
# e.g., RUN apk add --no-cache ca-certificates

# Expose the port your application listens on
EXPOSE 8080

# Start the application
CMD ["./main"]
