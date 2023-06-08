# Start from a lightweight and minimal base image
FROM golang:1.20-alpine3.18 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and cache Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application with optimizations
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o binary cmd/main.go

# Start a new stage using a minimal base image
FROM scratch

# Copy the built executable from the previous stage
COPY --from=builder /app/binary /app/binary
COPY --from=builder /app/config /app/config

# Set the working directory inside the container
WORKDIR /app

# Expose any necessary ports
EXPOSE 8000

# Run the Go application
CMD ["./binary"]
