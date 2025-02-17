#!/bin/bash

# Exit on error
set -e

# Add Go binary path to PATH
export PATH="$PATH:$(go env GOPATH)/bin"

# Install golint if not installed
if ! command -v golint &> /dev/null; then
    echo "Installing golint..."
    go install golang.org/x/lint/golint@latest
fi

echo "Running gofmt..."
# Find all .go files and run gofmt -w on them
find . -name "*.go" -not -path "./vendor/*" -exec gofmt -w {} \;

echo "Running golint..."
# Run golint on all packages
golint ./...

echo "All checks completed!"