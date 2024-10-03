#!/bin/bash

# Base path for the repository or vault
VAULT_PATH="/home/jwl/projects/projects/"

# Normalize the base path to ensure it's absolute and processed correctly
VAULT_PATH=$(realpath "$VAULT_PATH")

# Function to convert strings to kebab case
to_kebab_case() {
    echo "$1" | iconv -t ASCII//TRANSLIT | sed -r 's/[^a-zA-Z0-9]+/-/g' | sed -r 's/^-+\|-+$//g' | tr A-Z a-z
}

# Converts full file path to relative GitHub markdown link with URL encoding for spaces
make_link() {
    local path="$1"
    local dir="$2"
    local name=$(basename "$path")
    local cleanname="${name%.*}"
    local extension="${name##*.}"
    local kebab_name=$(to_kebab_case "$cleanname")
    local relative_path=$(realpath --relative-to="$dir" "$path")
    relative_path=$(to_kebab_case "$relative_path") # Apply kebab case to path
    relative_path="${relative_path// /%20}"         # Replace spaces with %20 for URL encoding

    if [[ "$extension" == "md" ]]; then
        # Markdown files lose their extension in the display
        echo "- [$kebab_name]($relative_path)"
    else
        # Non-Markdown files keep their extension in the display
        echo "- [$kebab_name.$extension]($relative_path)"
    fi
}

# Function to update a directory-named README file in each directory
update_readme() {
    local dir="$1"
    local base="$2"
    local dir_name=$(basename "README")
    local kebab_dir_name=$(to_kebab_case "$dir_name")
    local readme="${dir}/${kebab_dir_name}.md" # README named after the directory in kebab case

    echo "# Overview of $dir_name" >"$readme"
    echo "## Directories and Files" >>"$readme"
    echo "" >>"$readme"

    local total_files=0

    # Process files and subdirectories
    for item in "$dir"/*; do
        if [ -d "$item" ] && [[ "${item##*/}" != ".*" ]]; then
            local count=$(find "$item" -type f ! -name ".*" ! -name "${kebab_dir_name}.md" | wc -l)
            local item_name=$(basename "$item")
            local kebab_item_name=$(to_kebab_case "$item_name")
            echo "- Folder: $kebab_item_name - $count items" >>"$readme"
            total_files=$((total_files + count))
        elif [ -f "$item" ] && [[ "${item##*/}" != ".*" ]] && [[ "${item##*/}" != "${kebab_dir_name}.md" ]]; then
            make_link "$item" "$dir" >>"$readme"
            ((total_files++))
        fi
    done

    echo "" >>"$readme"
    echo "_${total_files} files and subdirectories._" >>"$readme"
}

# Export all directories excluding hidden ones and update their readme.md
find "$VAULT_PATH" -type d ! -path '*/.*' | while read -r directory; do
    update_readme "$directory" "$VAULT_PATH"
done
