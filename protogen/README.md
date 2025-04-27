# Model Context Protocol (MCP) Annotations

This package provides Protocol Buffer annotations for defining Model Context Protocol (MCP) services, methods, and fields. MCP is a protocol for defining and managing model context in AI applications, enabling consistent and well-documented APIs for model interactions.

## Service Level Annotations

### Basic Configuration
- `mcp_version` (string): Specifies the MCP protocol version for the service
- `mcp_type` (string): Defines the type of MCP service (e.g., "model", "context", "tool")

### Resource Configuration
- `mcp_resource_service` (bool): Indicates if the service is a resource service for model context management
- `mcp_resource_base_uri` (string): Base URI for resource operations in the model context
- `mcp_server_description` (string): Detailed description of the service's role in model context management

## Method Level Annotations

### Tool Configuration
- `mcp_tool_description` (string): Description of the tool/method's role in model context
- `mcp_tool_title` (string): Human-readable title for the tool in the model context
- `mcp_tool_read_only` (bool): Indicates if the tool only reads from the model context
- `mcp_tool_destructive` (bool): Indicates if the tool can modify the model context
- `mcp_tool_idempotent` (bool): Indicates if repeated calls with same arguments have no additional effect on the model context
- `mcp_tool_open_world` (bool): Indicates if the tool interacts with external entities outside the model context

## Field Level Annotations

### Parameter Types
- `mcp_path_param` (bool): Indicates if the field is a path parameter in the model context
- `mcp_query_param` (bool): Indicates if the field is a query parameter for model context operations
- `mcp_resource_uri` (bool): Indicates if the field represents a resource URI in the model context

### Field Properties
- `mcp_field_description` (string): Description of the field's role in the model context
- `mcp_required` (bool): Indicates if the field is required for model context operations
- `mcp_validation_pattern` (string): Regex pattern for validating field values in the model context
- `mcp_validation_message` (string): Custom validation error message for model context operations

## Message Level Annotations

### Resource Template Configuration
- `mcp_resource_template_message` (bool): Indicates if the message is a resource template for model context
- `mcp_uri_template` (string): URI template for the resource in the model context
- `mcp_template_description` (string): Description of the template's role in model context management

## Usage Example

```protobuf
syntax = "proto3";

package example.v1;

import "mcp/annotations.proto";

service ModelContextService {
  option (mcp.mcp_version) = "v1";
  option (mcp.mcp_type) = "model";
  option (mcp.mcp_resource_service) = true;
  option (mcp.mcp_resource_base_uri) = "/api/v1/model-context";
  option (mcp.mcp_server_description) = "Service for managing model context and interactions";

  rpc GetModelContext(GetModelContextRequest) returns (ModelContext) {
    option (mcp.mcp_tool_title) = "Get Model Context";
    option (mcp.mcp_tool_description) = "Retrieves the current model context by ID";
    option (mcp.mcp_tool_read_only) = true;
  }
}

message GetModelContextRequest {
  string context_id = 1 [(mcp.mcp_path_param) = true, (mcp.mcp_required) = true];
}
```

## Field Numbers

All MCP annotations use field numbers in the range 50000-50006 to avoid conflicts with other extensions. The field numbers are allocated as follows:

- Service options: 50000-50004
- Method options: 50000-50005
- Field options: 50000-50006
- Message options: 50000-50002 