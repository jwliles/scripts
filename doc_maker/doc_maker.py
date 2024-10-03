#!/usr/bin/env python3

import os

# Get the current directory
current_directory = os.getcwd()

# Loop through each file in the directory
for filename in os.listdir(current_directory):
    # Ignore directories and only process files
    if os.path.isfile(os.path.join(current_directory, filename)):
        # Get the file name without extension
        base_name = os.path.splitext(filename)[0]
        # Create a corresponding README filename
        readme_name = f"{base_name}.md"
        readme_path = os.path.join(current_directory, readme_name)

        # Create an empty Markdown file if it doesn't already exist
        if not os.path.exists(readme_path):
            with open(readme_path, 'w') as f:
                pass  # Just create an empty file
            print(f"Created: {readme_name}")
        else:
            print(f"Skipped: {readme_name} already exists")
