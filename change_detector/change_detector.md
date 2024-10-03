# Change Detector Script

## Overview

This Python script is designed to detect changes in files within a specified directory by comparing their hash values
against stored hashes. It leverages parallel processing to scan directories efficiently and logs various events,
including skipped files and errors.

## Features

- **File Hash Computation**: Calculates the hash for each file to detect modifications.
- **Change Detection**: Compares current file hashes with previously stored hashes to identify modified files.
- **Database Integration**: Saves file hashes and their modification times in a database for future comparisons.
- **Logging**: Logs information about skipped files, errors, and general progress during the scan.
- **Parallel Processing**: Utilizes a parallel scanning function for efficient file metadata retrieval.

## Usage

1. Ensure the necessary modules (`hash_manager`, `logger`) are available or installed, as they handle hash computation,
   storage, and logging functionality.
2. The `detect_changes` function is the main entry point. It requires:
    - `directory`: The path to the directory to be scanned.
    - `stored_hashes`: A dictionary containing previously stored hashes of files.
    - `scan_directory_with_parallelism`: A function that scans the directory in parallel and retrieves file metadata.
    - `DB_FILE`: The path to the database file for saving hashes and modification times.

3. Run the script:
   ```bash
   python3 change_detector.py
   ```

## Output

- **Changes List**: A list of file paths whose hashes have changed since the last scan.
- **Progress Indicator**: Displays the progress of the scan, including the number of processed files.
- **Log Files**: Errors and skipped files are logged for review.

## Requirements

- **Python 3**: The script is written in Python 3.
- **Logging Configuration**: The `logging` module is used for logging events and errors.
- **Custom Modules**: Ensure the presence of `hash_manager` and `logger` modules, which must contain:
    - `compute_file_hash(file_path)`: A function to calculate the file hash.
    - `save_file_hash(DB_FILE, file_path, file_hash, mtime)`: A function to store file hash information in a database.
    - `log_skipped_file(file_path, reason)`: A function to log files that are skipped during processing.
    - `log_event(event_type, message)`: A function to log general events.

## Customization

- **Adjust Progress Output**: You can modify the frequency of progress updates by changing the condition:
   ```python
   if files_processed % 100 == 0 or files_processed == total_files:
   ```

- **Modify Logging Behavior**: Customize how logging is handled in the `logger` module to suit your needs.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.