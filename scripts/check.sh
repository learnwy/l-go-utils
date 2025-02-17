#!/bin/bash

# Exit on error
set -e

# Run format checks
./scripts/format.sh

# Run linting
./scripts/lint.sh

# Run tests and coverage checks
./scripts/coverage.sh

echo "All checks completed!"