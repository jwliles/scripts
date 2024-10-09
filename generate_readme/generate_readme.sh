#!/bin/bash

# Base path for the repository or vault is set to the present working directory
VAULT_PATH=$(pwd)

# Normalize the base path to ensure it's absolute and processed correctly
VAULT_PATH=$(realpath "$VAULT_PATH")

# Current date in YYYY-MM-DD format
CURRENT_DATE=$(date "+%Y-%m-%d") # Ensure the date format is correct

# Function to convert names to snake_case
to_snake_case() {
    echo "$1" | iconv -t ASCII//TRANSLIT | sed -r 's/[^a-zA-Z0-9]+/_/g' | sed -r 's/^_+|_+$//g' | tr A-Z a-z
}

# Create a README.org file in each directory if it doesn't exist
create_readme_if_absent() {
    local dir="$1"
    local readme_path="$dir/README.org"

    # Check if the README.org file exists
    if [ ! -f "$readme_path" ]; then
        # Create an empty README.org
        touch "$readme_path"
        echo "Created $readme_path"
    fi
}

# Script to generate or update README.org for each directory, except the root
generate_readme() {
    local dir="$1"
    local dir_name=$(basename "$dir")
    local readme_path="$dir/README.org" # Always use README.org

    # Check if we are at the root directory to skip
    if [ "$dir" == "$VAULT_PATH" ]; then
        return # Skip the root directory README
    fi

    # Count the files excluding the README itself and hidden files
    local file_count=$(find "$dir" -maxdepth 1 -type f ! -name "README.org" ! -name ".*" | wc -l)

    # Header of the README
    echo "<!-- hash:$(md5sum <<<"$dir" | cut -d ' ' -f 1) -->" >"$readme_path"
    echo "* README" >>"$readme_path" # Always use "README" as the H1 heading
    echo "" >>"$readme_path"

    # Check and print the file count and current date
    if [ "$file_count" -gt 0 ]; then
        echo "> There are $file_count notes in this directory as of $CURRENT_DATE" >>"$readme_path"
    fi

    echo "" >>"$readme_path"
    echo "---" >>"$readme_path"
    echo "" >>"$readme_path"
    echo "** Categories" >>"$readme_path"
    echo "" >>"$readme_path"

    # Generate TOC for subdirectories, excluding hidden ones
    find "$dir" -maxdepth 1 -type d ! -name ".*" ! -path '*/.*' ! -samefile "$dir" | sort | while read subdir; do
        local subdir_name=$(basename "$subdir")
        local subdir_snake=$(to_snake_case "$subdir_name")
        echo "- [${subdir_name^}](*${subdir_snake})" >>"$readme_path"
    done

    echo "" >>"$readme_path"
    echo "---" >>"$readme_path"
    echo "" >>"$readme_path"

    # List notes in the current directory first
    echo "*** ${dir_name^}" >>"$readme_path"
    echo "" >>"$readme_path"

    # List files in the current directory, including the README.org, and excluding hidden ones
    find "$dir" -maxdepth 1 -type f ! -name ".*" ! -path '*/.*' | sort | while read file; do
        local file_name=$(basename "$file")
        local clean_name="${file_name%.*}"
        local relative_path="./$(realpath --relative-to="$dir" "$file" | sed 's/ /%20/g')"
        echo "- [$clean_name]($relative_path)" >>"$readme_path"
    done

    echo "" >>"$readme_path"

    # List each subdirectory separately after files
    find "$dir" -maxdepth 1 -type d ! -name ".*" ! -path '*/.*' ! -samefile "$dir" | sort | while read subdir; do
        local subdir_name=$(basename "$subdir")
        local subdir_snake=$(to_snake_case "$subdir_name")

        echo "*** ${subdir_name^}" >>"$readme_path" # Use the actual directory name as the heading
        echo "" >>"$readme_path"

        # List files within the subdirectory, excluding hidden ones
        find "$subdir" -maxdepth 1 -type f ! -name ".*" ! -path '*/.*' | sort | while read file; do
            local file_name=$(basename "$file")
            local clean_name="${file_name%.*}"
            local relative_path="./$(realpath --relative-to="$dir" "$file" | sed 's/ /%20/g')"
            echo "- [$clean_name]($relative_path)" >>"$readme_path"
        done

        echo "" >>"$readme_path"
    done
}

# Export the function and CURRENT_DATE to be available in subshells
export -f generate_readme to_snake_case create_readme_if_absent
export CURRENT_DATE

# Create README.org files if absent in all directories (including root)
find "$VAULT_PATH" -type d ! -name ".*" ! -path '*/.*' -exec bash -c 'create_readme_if_absent "$0"' {} \;

# Run the function on each directory except the root, excluding hidden directories
find "$VAULT_PATH" -mindepth 1 -type d ! -name ".*" ! -path '*/.*' -exec bash -c 'generate_readme "$0"' {} \;
