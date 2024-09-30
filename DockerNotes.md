### Prisma Migrations and Deployment Guide

This section outlines how Prisma migrations are handled in our Go project, the reasons behind the setup, and the key files involved.

#### **Why We Handle Migrations Separately**

In a production environment, running database migrations should be separate from the application build process. This ensures that:

- Migrations are applied to the target database just before or during deployment, ensuring consistency.
- The Docker build process remains clean, environment-agnostic, and doesn’t require access to the database.

By handling migrations as part of the deployment pipeline or during container startup, we ensure the application always runs against the latest database schema.

### Key Files Involved

1. **GitHub Workflow File (`.github/workflows/go.yml`)**:
   - This file defines the CI/CD pipeline for our project, handling tasks like building, testing, and running Prisma migrations before deploying our application.

2. **Dockerfile**:
   - The Dockerfile is responsible for building our Go application into a container image. It does not run migrations, keeping the image build process independent of database access.

### Step-by-Step Process

#### 1. GitHub Workflow (`.github/workflows/go.yml`)

The GitHub Actions workflow is configured to handle the following:

- **Checkout Code:** Retrieves the code from the repository.
- **Set Up Go Environment:** Configures the Go environment to build and test the application.
- **Install Prisma Client Go Dependency:** Downloads and installs the Prisma Client for Go.
- **Generate Prisma Client:** Uses the Prisma schema to generate the Go client.
- **Run Migrations:** Executes Prisma migrations to ensure the database schema is up-to-date.
- **Build and Test:** Builds the application and runs tests to verify everything works as expected.

Here’s the workflow configuration:

```yaml
name: Go

on:
  push:
    branches:
      - main
    paths-ignore:
      - "**/README.md"
  pull_request:
    branches:
      - main

permissions: read-all

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: 1.23.0

      - name: Download Required Dependencies
        run: go mod download

      - name: Add Prisma Client Go Dependency
        run: go get github.com/steebchen/prisma-client-go@latest

      - name: Generate Prisma Client
        run: go run github.com/steebchen/prisma-client-go generate

      - name: Run Prisma Migrations
        env:
          DATABASE_URL: ${{ secrets.DATABASE_URL }}
        run: go run github.com/steebchen/prisma-client-go migrate deploy

      - name: Build
        run: go build -v ./...

      - name: Test
        run: go test -v ./...
```

**Key Points:**

- The `DATABASE_URL` secret is configured in GitHub, allowing migrations to connect to the database.
- `go run github.com/steebchen/prisma-client-go migrate deploy` runs the migrations, ensuring the database is in sync with our schema.

#### 2. Dockerfile

The Dockerfile is focused purely on building the Go application without running migrations. This ensures the build process remains independent of any external services like a database.

```dockerfile
# Dockerfile

# Stage 1: Build the Go application
FROM golang:1.20-buster AS builder

WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

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
```

**Key Points:**

- The Dockerfile builds the Go application and doesn’t attempt to connect to or run migrations against the database.
- The resulting container image is clean, portable, and ready to run in any environment.

### Summary

- **Migrations are handled as part of the GitHub Actions workflow** to ensure they are executed against the correct database environment during deployment.
- The **Dockerfile builds the application** without requiring access to the database, keeping the build process environment-independent.

By separating these concerns, we maintain a clean, efficient, and reliable deployment process that ensures our application and database are always in sync.