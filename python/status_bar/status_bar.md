# Directory Scanning Script with Progress Bar

## Overview

This Python script scans through a specified directory and its subdirectories to collect metadata (file paths and
modification times) for all non-hidden files. It uses the `tqdm` library to display a progress bar, providing a visual
indication of the scan's progress.

## Features

- **Directory Traversal**: Recursively scans the specified directory and its subdirectories.
- **File Metadata Collection**: Collects file paths and modification times (`mtime`).
- **Progress Bar**: Uses `tqdm` to display a progress bar during the scan, showing how many files have been processed.

## Usage

1. **Install `tqdm`**: Ensure you have the `tqdm` library installed, which is used for the progress bar. You can install
   it with:
   ```bash
   pip install tqdm
   ```

2. **Run the Script**: Import the function `scan_directory_with_progress()` in your script, or call it directly within a
   `main()` function:
   ```python
   from scanner import scan_directory_with_progress

   directory = "/path/to/your/directory"
   files_metadata = scan_directory_with_progress(directory)
   ```

   Replace `"/path/to/your/directory"` with the path of the directory you wish to scan.

3. **Output**: The function returns a list of tuples containing:
    - **File Path**: The full path to each file.
    - **Modification Time**: The last modified time of each file (Unix timestamp).

## Example

```python
if __name__ == "__main__":
    directory = "/home/user/my_directory"
    metadata = scan_directory_with_progress(directory)

    for file_path, mtime in metadata:
        print(f"File: {file_path}, Last Modified: {mtime}")
```

This will print each file's path and its modification time, while showing a progress bar as the directory is being
scanned.

## Requirements

- **Python 3**: The script is written in Python 3.
- **`tqdm` Library**: For displaying the progress bar. Install it using:
  ```bash
  pip install tqdm
  ```

## Customization

- **Hidden File Exclusion**: The script ignores hidden files (those starting with `.`). You can modify the
  `if not file.startswith('.')` condition in the loop to include hidden files if desired.
- **Additional Metadata**: You can extend the `files_metadata` collection to include other file attributes like size (
  `os.path.getsize(file_path)`) or permissions (`os.stat(file_path).st_mode`).

## Notes

- **Progress Bar**: The progress bar will display the total number of files being processed. If the directory contains a
  large number of files, the progress bar provides real-time feedback on the script's progress.
- **Cross-Platform**: The script uses `os.walk` and `os.path` for file traversal and metadata collection, making it
  compatible with various operating systems (Linux, macOS, Windows).

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.