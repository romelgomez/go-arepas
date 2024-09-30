#!/bin/bash

# Check if Go is installed
if ! command -v go &> /dev/null
then
    echo "Go is not installed. Please install Go and try again."
    exit 1
fi

# Step 1: Download Go dependencies
echo "Downloading Go dependencies..."
go mod download

# Step 2: Generate the Prisma Client
echo "Generating Prisma Client..."
go run github.com/steebchen/prisma-client-go generate

# Step 3: Create .env file if it doesn't exist
if [ ! -f .env ]; then
    echo "Creating .env file from .env.example..."
    cp .env.example .env
    echo ".env file created. Please review and update it with the correct values."
else
    echo ".env file already exists. Skipping this step."
fi

# Step 4: Apply Prisma migrations for development
echo "Applying Prisma migrations..."
go run github.com/steebchen/prisma-client-go migrate dev

# Step 5: Build the Go application
echo "Building the application..."
go build -o main .

echo "Project bootstrap completed successfully!"
echo "You can start the application with: ./main"
