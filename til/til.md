# Directory-Named README Generator Script

## Overview

This Bash script generates or updates `README.md` files for each directory in a specified vault or repository path. Each
`README.md` file summarizes the contents of its corresponding directory, providing an overview of all files and
subdirectories. The script uses "kebab case" for naming and formatting and creates Markdown-friendly links.

## Features

- **Automatic README Creation**: Creates a `README.md` file in each directory, named after the directory itself in kebab
  case.
- **Directory and File Listing**: Lists all files and subdirectories in the `README.md`, providing a count of items in
  each subdirectory and linking to each file.
- **URL Encoding for Links**: Converts spaces to `%20` in URLs for compatibility with GitHub Markdown.
- **Kebab Case Naming**: Converts file and directory names to kebab case to maintain consistent naming conventions.

## Usage

1. **Set the Vault/Repository Path**: Ensure the `VAULT_PATH` variable in the script points to the base directory where
   you want the `README.md` files to be created/updated:
   ```bash
   VAULT_PATH="/path/to/your/vault"
   ```

2. **Run the Script**: Make the script executable and then execute it:
   ```bash
   chmod +x update_readmes.sh
   ./update_readmes.sh
   ```

   The script will traverse all non-hidden directories within `VAULT_PATH`, creating or updating a `README.md` file in
   each one.

3. **Output**: Each `README.md` will:
    - Have a header titled "Overview of `<directory-name>`".
    - List all subdirectories with a count of items within them.
    - Include kebab case links to each file and subdirectory, with appropriate formatting for Markdown.

## Example

**Directory Structure**:

```
/projects/
    /my-first-project/
        script.sh
        README.md
    /another_project/
        data.txt
        config.yml
```

After running the script, `README.md` files will be generated for each directory:

- **`projects/my-first-project/README.md`**:
  ```markdown
  # Overview of my-first-project
  ## Directories and Files

  - [script.sh](script.sh)

  _1 files and subdirectories._
  ```

- **`projects/another-project/README.md`**:
  ```markdown
  # Overview of another-project
  ## Directories and Files

  - [data.txt](data.txt)
  - [config.yml](config.yml)

  _2 files and subdirectories._
  ```

## Requirements

- **Bash Shell**: The script is designed to run in a Bash shell.
- **Read and Write Permissions**: Ensure the script has the necessary permissions to read from the source directory and
  write to each target directory.
- **Standard Utilities**: The script uses common Linux utilities like `realpath`, `sed`, `find`, and `iconv`.

## Customization

- **Exclude Files/Directories**: By default, hidden files (names starting with `.`) and `README.md` files are excluded.
  You can customize this behavior by modifying the conditions in the `update_readme()` function.
- **Kebab Case Conversion**: If you wish to use a different naming convention, modify the `to_kebab_case()` function as
  necessary.

## Notes

- **Directory Naming in README.md**: Each `README.md` is named after its parent directory in kebab case.
- **Non-Markdown Files**: If a file is not a Markdown file (`.md`), its extension is included in the display name;
  otherwise, the extension is removed.
- **Script Recursion**: The script recursively traverses all directories from `VAULT_PATH` but skips hidden directories
  and files.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.