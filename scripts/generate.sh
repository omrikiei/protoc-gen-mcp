#!/bin/bash

set -e

# Create output directories
mkdir -p internal/protogen/mcp

# Generate Go code from proto files
protoc \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  internal/protogen/mcp/annotations.proto \
  internal/protogen/mcp/service.proto

# Generate Go code for the example
protoc \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  --mcp_out=. \
  --mcp_opt=paths=source_relative \
  examples/helloworld/helloworld.proto \
  examples/complex/complex.proto 