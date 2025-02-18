#!/bin/bash

# Exit on error
set -e

# Add Go binary path to PATH
export PATH="$PATH:$(go env GOPATH)/bin"

echo "Running tests with coverage..."
# Create coverage output directory if it doesn't exist
mkdir -p coverage

# Run tests with coverage, excluding examples directory
go test $(go list ./... | grep -v /examples/) -coverprofile=coverage/coverage.out

# Generate HTML coverage report
go tool cover -html=coverage/coverage.out -o coverage/coverage.html

# Generate text coverage report
go tool cover -func=coverage/coverage.out -o coverage/coverage.txt

# Check coverage percentage
COVERAGE=$(go tool cover -func=coverage/coverage.out | grep total: | awk '{print $3}' | sed 's/%//')
MIN_COVERAGE=80.0

echo "Code coverage: $COVERAGE%"
if (( $(echo "$COVERAGE < $MIN_COVERAGE" | bc -l) )); then
    echo "Error: Code coverage is below $MIN_COVERAGE% threshold"
    exit 1
fi

echo "Coverage check completed!"