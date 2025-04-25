.PHONY: all build install clean protogen generate

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=protoc-gen-mcp
INSTALL_PATH=$(GOPATH)/bin

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) cmd/protoc-gen-mcp/main.go

install: build
	mkdir -p $(INSTALL_PATH)
	cp $(BINARY_NAME) $(INSTALL_PATH)/$(BINARY_NAME)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(INSTALL_PATH)/$(BINARY_NAME)

# Generate protobuf files for internal/protogen
protogen:
	mkdir -p protogen/mcp
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		protogen/mcp/annotations.proto \
		protogen/mcp/service.proto

# Generate example protobuf files
examples:
	go build -o ~/go/bin/protoc-gen-mcp cmd/protoc-gen-mcp/main.go && ./scripts/generate.sh

# Generate all protobuf files (both internal and examples)
generate: protogen examples
	@echo "All protobuf files generated successfully"
