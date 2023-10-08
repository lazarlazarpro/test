# Use the official Go image from DockerHub as the base image
FROM golang:1.17

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN cd /app/main && go build -o main

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["/app/main/main"]
