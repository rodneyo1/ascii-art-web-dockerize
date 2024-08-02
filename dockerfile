# Use a lightweight base image for Go applications
FROM golang:alpine AS builder

# Set the working directory for the build stage
WORKDIR /app

# Copy the project code to the build stage
COPY . .

# Install dependencies during build
RUN go mod download

# Build the Go binary
RUN go build -o asciiartserver main.go

# Use a slimmer image for the final container
FROM alpine:latest

# Copy the binary from the build stage
COPY --from=builder /app/asciiartserver /app/asciiartserver

# Set the working directory for the final container
WORKDIR /app

# Expose the port where the server will listen
EXPOSE 8080

# Command to run the server
CMD ["asciiartserver"]
