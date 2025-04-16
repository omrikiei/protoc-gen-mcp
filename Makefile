.PHONY: all build install clean

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

# Example usage target
example:
	cd examples/helloworld && \
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		--mcp_out=. --mcp_opt=paths=source_relative \
		helloworld.proto 
generate:
	go build -o ~/go/bin/protoc-gen-mcp cmd/protoc-gen-mcp/main.go && ./scripts/generate.sh
