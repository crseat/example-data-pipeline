# Use the official Golang image as a build stage
FROM golang:1.19 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o /main main.go

# Use a smaller base image for the final image
FROM alpine:latest

# Set environment variables
ENV APP_ENV=production
ENV SERVER_PORT=:8080
ENV KAFKA_BROKER=kafka:9092
ENV KAFKA_TOPIC=post-topic

# Copy the Pre-built binary file from the builder stage
COPY --from=builder /main /main

# Expose port
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["/main"]
