# Delete Empty Markdown Files Script

## Overview

This Python script scans through a directory and deletes Markdown files that are considered "empty." A file is
considered empty if it only contains YAML frontmatter and possibly an H1 heading, with no other content.

## Features

- **Empty File Detection**: Identifies Markdown files (`.md`) that only contain:
    - YAML frontmatter (typically found at the top of Markdown files between `---` lines).
    - An optional H1 heading (`# Heading`).
- **Deletion of Empty Files**: Deletes detected empty Markdown files to help maintain a clean directory.
- **Dry Run Option**: Allows a preview of files that would be deleted without actually removing them.

## Usage

1. **Run the Script**: Use the following command to execute the script:
   ```bash
   python3 delete-empty-files.py
   ```
   By default, it will scan the current working directory.

2. **Dry Run Option**: To perform a dry run without deleting any files, add the `--dry-run` flag:
   ```bash
   python3 delete-empty-files.py --dry-run
   ```
   This will list all the files that would be deleted.

## Requirements

- **Python 3**: The script is written in Python 3.
- **Markdown Files**: The script targets files with a `.md` extension.
- **YAML and H1 Detection**: Uses regular expressions to detect the presence of YAML frontmatter and an optional H1
  heading.

## Customization

- **Directory Scanning**: By default, the script scans the current working directory. You can modify the
  `current_directory` variable or adapt the `find_and_delete_empty_files()` function to scan a different directory.
- **Criteria for Empty Files**: The script uses regex patterns to define what constitutes an empty file. You can adjust
  these patterns (`yaml_pattern` and `h1_pattern`) within the `is_empty_file()` function as needed.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.