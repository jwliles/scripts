# Index Generation Script

## Overview

This Bash script automates the creation and updating of `README.md` files in each directory of a vault or repository,
except the root directory. It provides a file count, lists categories, generates a Table of Contents (TOC) for
subdirectories, and organizes notes or files for better navigation within the vault.

## Features

- **Creates `README.md` Files**: Generates a `README.md` file in each directory if it doesn't exist.
- **Updates `README.md` Files**: Populates each `README.md` with:
    - A unique hash as a comment (for tracking changes).
    - A header with the directory name.
    - A count of the files in the directory.
    - A Table of Contents for subdirectories.
    - A list of files in each directory and its subdirectories.
- **Excludes Hidden Files/Directories**: Ignores files or directories that are hidden (names starting with `.`).

## Usage

1. **Navigate to the Vault/Repository Directory**:
   ```bash
   cd /path/to/your/vault
   ```

2. **Run the Script**:
   ```bash
   ./generate_readmes.sh
   ```

   The script will:
    - Traverse through each directory in the vault.
    - Create or update a `README.md` file in each directory.
    - Skip hidden files and directories.

3. **Output**:
    - **`README.md` Files Created**: A new `README.md` will be created in directories that don't have one.
    - **`README.md` Files Updated**: Existing `README.md` files will be updated with a list of notes/files and a TOC for
      any subdirectories.

## Requirements

- **Bash**: The script is intended for a Bash shell.
- **`realpath` Utility**: Used to resolve and normalize directory paths.
- **`find`, `sed`, `iconv`, `md5sum` Utilities**: Standard command-line tools used for directory traversal, string
  manipulation, and hashing.

## Customization

- **Modify `VAULT_PATH`**: The base directory for processing is set to the current working directory. You can modify
  this by changing the `VAULT_PATH` variable:
  ```bash
  VAULT_PATH="/your/desired/path"
  ```

- **File Naming and TOC Generation**: You can adjust the TOC generation and the file naming convention by tweaking the
  `generate_readme` and `to_snake_case` functions.

## Notes

- **`README.md` Structure**: Each `README.md` file contains:
    - A header with the directory name and a count of notes/files.
    - A list of all files in the directory, excluding hidden files.
    - A categorized TOC listing any subdirectories and their notes.
- **Skipping Root Directory**: The script does not generate a `README.md` for the root directory.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.