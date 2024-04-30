# Use a base image with Go 1.19
FROM golang:1.19 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go Mod files first (to leverage Docker cache)
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application without CGO for Linux
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o grpcServer .

# Use a minimal base image to run the application
FROM alpine:latest
WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/grpcServer .

# Expose the port 

# Command to run the executable
CMD ["./grpcServer"]