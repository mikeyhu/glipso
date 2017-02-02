#!/usr/bin/env bash
set -e

echo "Formatting..."
go fmt ./...

echo "Tests..."
go test ./...

echo "Building..."
go build

echo "Successfully built mekkanism!"