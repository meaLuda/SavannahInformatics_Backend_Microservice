.PHONY: all build docker-build docker-run clean

# Define variables for the binary and image names
BINARY_NAME=ordersapi
IMAGE_NAME=ordersapi

# Default target: build the binary and Docker image
docker-starter: docker-build docker-run

# Build the Go binary
build:
	@echo "Building the binary..."
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) main.go

# Build the Docker image
docker-build:
	@echo "Building the Docker image..."
	docker build -t $(IMAGE_NAME) .

# Run the Docker container
docker-run:
	@echo "Running the Docker container..."
	docker run -d -p 8080:8080 $(IMAGE_NAME)

# Clean up build artifacts
docker-clean:
	@echo "Cleaning up..."
	docker rm -v $(IMAGE_NAME)

