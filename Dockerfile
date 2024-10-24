# Start from the official Go image
FROM golang:1.21 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o myapp

# Use a lightweight base image for the final stage
FROM alpine:3.18

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/myapp .

# Expose port if necessary (for example, if your app listens on port 8080)
EXPOSE 8080

# Run the application
CMD ["./myapp"]