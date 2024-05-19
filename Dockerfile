FROM golang:1.22.3-alpine3.18

WORKDIR /app

# Copy the entire project to the working directory
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# # Install the package
RUN go install -v ./...

# Create a new file .env
RUN touch .env

# Build the Go binary
RUN go build -o go-api-auth

# Expose port 8080 to the outside world
EXPOSE 8888

# Command to run the executable
CMD ./go-api-auth
