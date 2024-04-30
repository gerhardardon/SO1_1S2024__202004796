# Start from the golang base image
FROM golang:1.19 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files if they exist
COPY go.mod go.sum* ./

# Download all dependencies if go mod and sum files exist
RUN if [ -f go.mod ]; then go mod download; fi

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o grpcClient .

# Start from a fresh alpine image to create a small final image
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the builder stage
COPY --from=builder /app/grpcClient .

# No default EXPOSE directive is needed here unless the client also serves as a server
EXPOSE 3000
# Run the binary program produced by `go install`
CMD ["./grpcClient"]