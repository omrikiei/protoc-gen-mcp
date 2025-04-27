# protoc-gen-mcp

A Protocol Buffers compiler plugin that generates MCP (Management Control Plane) server, client, and mock server code from protobuf service definitions.

## Overview

This project is a Protocol Buffers compiler plugin that generates:
- MCP server implementation
- MCP client implementation
- Mock server for testing

The generated code follows the MCP specification and can be used to implement management control plane functionality in your services.

## Installation

```bash
go install github.com/omrikiei/protoc-gen-mcp/cmd/protoc-gen-mcp
```

## Usage

1. Add the plugin to your protoc command:

```bash
protoc --mcp_out=. --mcp_opt=paths=source_relative your_service.proto
```

2. Or use with buf:

```yaml
# buf.gen.yaml
version: v1
plugins:
  - plugin: mcp
    path: protoc-gen-mcp
    out: .
    opt:
      - paths=source_relative
```

## MCP Annotations

The plugin supports various annotations to configure MCP services, methods, and fields. Here's how to use them in your `.proto` files:

### Service Level Annotations

```protobuf
service YourService {
  option (mcp.annotations.mcp_service) = true;  // Required to enable MCP generation
  option (mcp.annotations.mcp_version) = "v1";  // API version
  option (mcp.annotations.mcp_type) = "your-type";  // Service type
  option (mcp.annotations.mcp_resource_service) = true;  // Enable resource handling
  option (mcp.annotations.mcp_resource_base_uri) = "your://";  // Base URI for resources
  option (mcp.annotations.mcp_server_description) = "Description of your service";
}
```

### Method Level Annotations

```protobuf
rpc YourMethod (Request) returns (Response) {
  option (mcp.annotations.mcp_tool) = true;  // Required to enable MCP tool generation
  option (mcp.annotations.mcp_tool_description) = "Description of the tool";
  option (mcp.annotations.mcp_tool_title) = "Human readable title";
  option (mcp.annotations.mcp_tool_read_only) = true;  // Tool doesn't modify environment
  option (mcp.annotations.mcp_tool_destructive) = false;  // Tool may perform destructive updates
  option (mcp.annotations.mcp_tool_idempotent) = true;  // Repeated calls have no additional effect
  option (mcp.annotations.mcp_tool_open_world) = false;  // Tool interacts with external entities
}
```

### Field Level Annotations

```protobuf
message YourMessage {
  string field1 = 1 [
    (mcp.annotations.mcp_path_param) = true,  // Field is a path parameter
    (mcp.annotations.mcp_query_param) = true,  // Field is a query parameter
    (mcp.annotations.mcp_resource_uri) = true,  // Field is a resource URI
    (mcp.annotations.mcp_field_description) = "Description of the field",
    (mcp.annotations.mcp_required) = true,  // Field is required
    (mcp.annotations.mcp_validation_pattern) = "^[a-zA-Z0-9_]{3,50}$",  // Validation regex
    (mcp.annotations.mcp_validation_message) = "Validation error message"
  ];
}
```

### Message Level Annotations

```protobuf
message YourMessage {
  option (mcp.annotations.mcp_resource_template_message) = true;  // Enable resource template
  option (mcp.annotations.mcp_uri_template) = "your://{field1}/{field2}";  // URI template
  option (mcp.annotations.mcp_template_description) = "Description of the template";
}
```

## Buf Configuration

To use the plugin with buf, you need to configure both `buf.gen.yaml` and `buf.work.yaml`:

### buf.gen.yaml

```yaml
version: v1
plugins:
  - name: mcp
    out: gen/go
    opt:
      - paths=source_relative
  - name: go
    out: gen/go
    opt:
      - paths=source_relative
  - name: go-grpc
    out: gen/go
    opt:
      - paths=source_relative
```

### buf.work.yaml

```yaml
version: v1
directories:
  - proto
  - internal/protogen/mcp
```

### buf.yaml

```yaml
version: v1
name: buf.build/omrikiei/protoc-gen-mcp
deps:
  - buf.build/googleapis/googleapis
  - buf.build/grpc-ecosystem/grpc-gateway
```

## Project Structure

```
protoc-gen-mcp/
├── cmd/
│   └── protoc-gen-mcp/     # Main plugin entry point
├── internal/
│   ├── generator/          # Code generation logic
│   ├── protogen/          # Protocol definitions
│   └── template/          # Code templates
├── pkg/                   # Public packages
└── examples/              # Example usage
```

## Development

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Build the plugin:
   ```bash
   go build -o protoc-gen-mcp ./cmd/protoc-gen-mcp
   ```

## License

MIT License 

# MCP Server Example

This is an example server implementation using the MCP (Microservice Communication Protocol) framework. It demonstrates various features including path parameters, query parameters, health checks, and metrics.

## Features

- Path parameter handling
- Query parameter handling
- Health checks (`/healthz` and `/readyz`)
- Prometheus metrics (`/metrics`)
- Graceful shutdown
- Structured logging
- Docker support

## Endpoints

### Greeter Service

- `GET /v1/hello/{name}` - Simple greeting with path parameter
- `GET /v1/hello?name=value&lang=en&formal=true&repeat=3` - Greeting with query parameters

### Complex Service

- `GET /v1/users/{user_id}/profiles/{profile_id}` - Get user profile with multiple path parameters
- `GET /v1/orgs/{org_id}/teams/{team_id}/members/{member_id}` - Get nested resource
- `GET /v1/resources?query=search&page=1&page_size=10` - Search resources with complex query parameters

### Health and Metrics

- `GET /healthz` - Health check endpoint
- `GET /readyz` - Readiness check endpoint
- `GET /metrics` - Prometheus metrics endpoint

## Running the Server

### Local Development

1. Install dependencies:
```bash
go mod download
```

2. Generate protobuf code:
```bash
./scripts/generate.sh
```

3. Build and run the server:
```bash
go build -o server cmd/server/main.go
./server -port 8080 -metrics-port 9090
```

### Docker

1. Build and run using Docker Compose:
```bash
docker-compose up --build
```

2. Or build and run manually:
```bash
docker build -t mcp-server .
docker run -p 8080:8080 -p 9090:9090 mcp-server
```

## Testing the Server

### Greeter Service

```bash
# Path parameter
curl http://localhost:8080/v1/hello/john

# Query parameters
curl "http://localhost:8080/v1/hello?name=john&lang=en&formal=true&repeat=3"
```

### Complex Service

```bash
# User profile
curl http://localhost:8080/v1/users/123/profiles/456

# Nested resource
curl http://localhost:8080/v1/orgs/org1/teams/team1/members/member1

# Search resources
curl "http://localhost:8080/v1/resources?query=test&page=1&page_size=10&filters=type:doc&sort_by=name&sort_desc=true"
```

### Health and Metrics

```bash
# Health check
curl http://localhost:8080/healthz

# Readiness check
curl http://localhost:8080/readyz

# Metrics
curl http://localhost:9090/metrics
```

## Configuration

The server can be configured using command-line flags:

- `-port`: Port to listen on for the main server (default: 8080)
- `-metrics-port`: Port to listen on for metrics (default: 9090)

## Monitoring

The server exposes Prometheus metrics that can be scraped by a monitoring system. Key metrics include:

- `http_requests_total`: Total number of HTTP requests
- `http_request_duration_seconds`: HTTP request duration

## Development

### Adding New Services

1. Create a new proto file in the `examples` directory
2. Add the service definition with MCP annotations
3. Update the generate script to include the new proto file
4. Implement the service in the server code

### Testing

Run the tests:
```bash
go test ./...
```

## License

MIT 
