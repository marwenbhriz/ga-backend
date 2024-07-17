# Start from a base Go image
ARG GO_VERSION=1.22.4

FROM golang:${GO_VERSION} builder

# Set necessary environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main ./

# Expose port 9090 to the outside world
EXPOSE 9090

# Command to run the executable
CMD ["./main"]