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

echo "Running golint..."
# Run golint on all packages
golint ./...

echo "Lint check completed!"