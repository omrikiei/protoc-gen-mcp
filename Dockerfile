# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Generate protobuf code
RUN go install github.com/omrikiei/protoc-gen-mcp/cmd/protoc-gen-mcp
RUN protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --mcp_out=. --mcp_opt=paths=source_relative \
    examples/helloworld/helloworld.proto

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/server/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/server .

# Expose ports
EXPOSE 8080 9090

# Run the server
CMD ["./server"] 