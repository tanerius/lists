# Base image for running tests
FROM golang:1.23.0 AS tester

# Set the working directory
WORKDIR /app

# Copy everything from the builder
COPY . .

# Default command for tests
CMD ["go", "test", "./..."]