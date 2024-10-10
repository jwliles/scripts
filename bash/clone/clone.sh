#!/bin/bash

# Check if a target directory is provided
if [ -z "$1" ]; then
    echo "Usage: $0 <target-directory>"
    exit 1
fi

# Set the target directory
TARGET_DIR="$1"

# Create the target directory if it doesn't exist
mkdir -p "$TARGET_DIR"

# Loop over all Git repositories in the current directory
for repo in */.git; do
    # Extract the project name from the path
    PROJECT_NAME=$(basename "$(dirname "$repo")")

    # Clone the repository to the target directory
    echo "Cloning $PROJECT_NAME to $TARGET_DIR/$PROJECT_NAME"
    git clone "$(dirname "$repo")" "$TARGET_DIR/$PROJECT_NAME"
done

echo "All repositories cloned to $TARGET_DIR"
