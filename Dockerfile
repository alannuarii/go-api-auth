FROM golang:1.22.3-alpine3.18

WORKDIR /app

# Copy the entire project to the working directory
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# # Install the package
RUN go install -v ./...

# Build the Go binary
RUN go build -o main.go

# Expose port 8080 to the outside world
EXPOSE 8888

# Command to run the executable
CMD ["go", "run", "main.go"]
