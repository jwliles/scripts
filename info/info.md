# Directory Information Script

## Overview

This Bash script provides an overview of the contents of the current directory. It reports the total size of the
directory, the number of files and folders, and identifies any empty files or directories.

## Features

- **Current Directory Information**: Uses the current directory as the target for analysis.
- **Total Size Calculation**: Displays the total size of the directory in gigabytes.
- **File and Folder Count**: Counts and displays the total number of files and folders.
- **Empty File and Directory Identification**: Provides a count of any empty files or directories within the current
  directory.

## Usage

1. Save the script to a file, e.g., `directory_info.sh`.
2. Make the script executable:
   ```bash
   chmod +x directory_info.sh
   ```
3. Run the script from the desired directory:
   ```bash
   ./directory_info.sh
   ```

## Output Details

- **Directory Information for**: Displays the current directory being analyzed.
- **Total Size (in GB)**: The total size of the directory and its contents in gigabytes.
- **Total Files**: The count of all files (including those in subdirectories).
- **Total Folders**: The count of all folders (including subdirectories).
- **Total Empty Files**: The count of files that are empty (0 bytes in size).
- **Total Empty Directories**: The count of directories that do not contain any files or subdirectories.

## Requirements

- **Bash**: The script should be run in a Bash shell.
- **`du`, `find`, `awk`, and `wc` utilities**: These standard command-line tools are used for directory size
  calculations, file/folder counting, and filtering.

## Customization

You can modify the `DIR` variable to point to any specific directory if you don't want to use the current directory:

```bash
DIR="/path/to/your/directory"
```

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.