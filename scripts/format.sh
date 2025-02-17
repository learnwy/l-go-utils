#!/bin/bash

# Exit on error
set -e

# Add Go binary path to PATH
export PATH="$PATH:$(go env GOPATH)/bin"

echo "Running gofmt..."
# Find all .go files and run gofmt -w on them
find . -name "*.go" -not -path "./vendor/*" -exec gofmt -w {} \;

echo "Format check completed!"