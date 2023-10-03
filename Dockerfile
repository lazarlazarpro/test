# Use the official Go image from DockerHub as the base image
FROM golang:1.17

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

# Change to the directory containing main.go
WORKDIR /app/main

# Build the Go app
RUN go build -o main

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
