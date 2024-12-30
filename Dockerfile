# Use the official Go image as a base
FROM golang:1.20 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the application source code to the container
COPY . .

# Build the application
RUN go build -o todocli main.go

# Use a minimal base image for the final executable
FROM debian:bullseye-slim

# Set the working directory for the runtime
WORKDIR /app

# Copy the compiled binary from the builder
COPY --from=builder /app/todocli .

# Expose ports (if your app uses networking, not CLI only)
# EXPOSE 8080

# Set the entry point for the container
ENTRYPOINT ["./todocli"]
