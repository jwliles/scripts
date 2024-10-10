# Script Index Generator

## Overview

This Python script scans a directory for various script files (such as `.sh`, `.py`, `.go`, `.rs`, `.pl`, `.rb`) and
generates a `README.md` file containing basic information about each script. The README file includes the file path,
shebang, and extracted single-line comments to provide a description for each script.

## Features

- **File Type Support**: Automatically detects and processes files with extensions `.sh`, `.py`, `.go`, `.rs`, `.pl`,
  and `.rb`.
- **Shebang Detection**: Extracts and displays the shebang (`#!`) from each script to indicate its interpreter.
- **Description Extraction**: Collects single-line comments from scripts to provide a brief description or notes about
  the functionality.
- **Non-Recursive**: Scans only the root directory without traversing into subdirectories.

## Usage

1. **Place the Script in the Target Directory**: Save this script in the directory containing the scripts you want to
   index.

2. **Run the Script**:
   ```bash
   python3 generate_script_readme.py
   ```

   The script will:
    - Scan the current directory for script files.
    - Extract metadata such as the shebang and comments from each file.
    - Generate a `README.md` file containing an index of all detected scripts.

3. **Output**:
    - A `README.md` file will be created in the same directory, with sections for each script that include:
        - **File Name**: The script's name.
        - **Path**: The path to the script file.
        - **Shebang**: The shebang line to indicate the interpreter.
        - **Description**: Comments extracted from the script, formatted for easy reading.

## Requirements

- **Python 3**: The script is written in Python 3.
- **Script Files**: The script targets files ending in `.sh`, `.py`, `.go`, `.rs`, `.pl`, and `.rb`. You can modify the
  extensions in the code to support other file types.

## Customization

- **Target Directory**: By default, the script scans the current directory (`"."`). To change this, modify the
  `generate_readme` function to point to a specific path:
  ```python
  generate_readme("/path/to/your/scripts")
  ```

- **File Extensions**: To add or remove file types, edit the list of supported extensions in the `generate_readme`
  function:
  ```python
  if os.path.isfile(file_path) and file.endswith((".sh", ".py", ".your_extension")):
  ```

## Notes

- **Single-Line Comments Only**: The script only extracts single-line comments starting with `#` or `//`.
- **Inline Code Formatting**: Comments are wrapped in backticks to appear as inline code in the generated README.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.