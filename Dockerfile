FROM golang:1.20

WORKDIR /usr/app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY *.go ./

# Run tests
CMD ["go", "test", "./..."]