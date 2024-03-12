#!/bin/bash

# Navigate to the script's directory (project root directory)
cd "$(dirname "$0")"

# Find all subdirectories containing .go files and build executables
find . -type f -name '*.go' ! -path "./vendor/*" -execdir go build \;

echo "Build complete."
