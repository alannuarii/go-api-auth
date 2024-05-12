FROM golang:1.22.3-alpine3.18

WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the Go binary
RUN go build -o go-api-auth

# Expose port 8080 to the outside world
EXPOSE 8888

# Command to run the executable
CMD ["./go-api-auth"]
