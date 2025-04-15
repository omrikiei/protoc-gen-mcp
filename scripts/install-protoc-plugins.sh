#!/bin/bash

set -e

# Install protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Install protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest 