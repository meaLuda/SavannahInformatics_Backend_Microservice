# Step 1: Build the Go binary
FROM golang:1.22.1-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code
COPY . .
COPY .env .

# Copy go mod and sum files
RUN go mod tidy

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /echo-api

# Step 2: Create a small final image
FROM scratch


COPY .env .
# Copy the binary from the builder stage
COPY --from=builder /echo-api /echo-api

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/echo-api"]
