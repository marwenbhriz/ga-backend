# Start from a base Go image
FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Expose port 9090 to the outside world
EXPOSE 9090

# Command to run the application using go run
CMD ["go", "run", "main.go"]