# README File Generator Script

## Overview

This Python script scans the current directory and creates an empty Markdown README file (`.md`) for each file that does
not already have a corresponding README. It ensures that each file in the directory has a README file with the same base
name.

## Features

- **Automatic README Creation**: For each file in the directory, an empty Markdown file (`.md`) is created if it doesn't
  already exist.
- **File Matching by Name**: Each README is named after the corresponding file, with the same base name but with a `.md`
  extension.
- **Skip Existing README Files**: The script checks for existing README files and skips their creation if they are
  already present.

## Usage

1. **Run the Script**: Navigate to the desired directory and execute the script:
   ```bash
   python3 create_readmes.py
   ```
   This will generate `.md` files for each file without an existing README in the current directory.

2. **Output**:
    - **Created**: If a README file is created, you will see a message like:
      ```
      Created: <filename>.md
      ```
    - **Skipped**: If a README file already exists, you will see a message like:
      ```
      Skipped: <filename>.md already exists
      ```

## Requirements

- **Python 3**: The script is written in Python 3.
- **Directory Scanning**: The script scans the current working directory for all files and creates README files
  accordingly.

## Customization

- **Target Directory**: By default, the script works in the current directory. To change this behavior, modify
  `current_directory` to point to a specific path:
  ```python
  current_directory = "/path/to/your/directory"
  ```

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.