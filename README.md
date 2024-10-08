# README

# Change Detector Script

## Overview

This Python script is designed to detect changes in files within
a specified directory by comparing their hash values against
stored hashes. It leverages parallel processing to scan
directories efficiently and logs various events,
including skipped files and errors.

## Features

- **File Hash Computation**: Calculates the hash for each file
  to detect modifications.
- **Change Detection**: Compares current file hashes with
  previously stored hashes to identify modified files.
- **Database Integration**: Saves file hashes and their
  modification times in a database for future comparisons.
- **Logging**: Logs information about skipped files, errors, and
  general progress during the scan.
- **Parallel Processing**: Utilizes a parallel scanning function
  for efficient file metadata retrieval.

## Usage

1. Ensure the necessary modules (`hash_manager`, `logger`) are
   available or installed, as they handle hash computation,
   storage, and logging functionality.

2. The `detect_changes` function is the main entry point. It
   requires:
   
   - `directory`: The path to the directory to be scanned.
   - `stored_hashes`: A dictionary containing previously stored
     hashes of files.
   - `scan_directory_with_parallelism`: A function that scans
     the directory in parallel and retrieves file metadata.
   - `DB_FILE`: The path to the database file for saving hashes
     and modification times.

3. Run the script:
   
   ```bash
   python3 change_detector.py
   ```

## Output

- **Changes List**: A list of file paths whose hashes have
  changed since the last scan.
- **Progress Indicator**: Displays the progress of the scan,
  including the number of processed files.
- **Log Files**: Errors and skipped files are logged for review.

## Requirements

- **Python 3**: The script is written in Python 3.
- **Logging Configuration**: The `logging` module is used for
  logging events and errors.
- **Custom Modules**: Ensure the presence of `hash_manager` and
  `logger` modules, which must contain:
  - `compute_file_hash(file_path)`: A function to calculate
    the file hash.
  - `save_file_hash(DB_FILE, file_path, file_hash, mtime)`: A
    function to store file hash information in a database.
  - `log_skipped_file(file_path, reason)`: A function to log
    files that are skipped during processing.
  - `log_event(event_type, message)`: A function to log
    general events.

## Customization

- **Adjust Progress Output**: You can modify the frequency of
  progress updates by changing the condition:
  
  ```python
  if files_processed % 100 == 0 or files_processed == total_files:
  ```

- **Modify Logging Behavior**: Customize how logging is handled
  in the `logger` module to suit your needs.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.

# delete-empty-files

# Delete Empty Markdown Files Script

## Overview

This Python script scans through a directory and deletes
Markdown files that are considered "empty." A file is considered
empty if it only contains YAML frontmatter and possibly an H1
heading, with no other content.

## Features

- **Empty File Detection**: Identifies Markdown files (`.md`)
  that only contain:
  - YAML frontmatter (typically found at the top of Markdown
    files between `---` lines).
  - An optional H1 heading (`# Heading`).
- **Deletion of Empty Files**: Deletes detected empty Markdown
  files to help maintain a clean directory.
- **Dry Run Option**: Allows a preview of files that would be
  deleted without actually removing them.

## Usage

1. **Run the Script**: Use the following command to execute the
   script:
   
   ```bash
   python3 delete-empty-files.py
   ```
   
   By default, it will scan the current working directory.

2. **Dry Run Option**: To perform a dry run without deleting any
   files, add the `--dry-run` flag:
   
   ```bash
   python3 delete-empty-files.py --dry-run
   ```
   
   This will list all the files that would be deleted.

## Requirements

- **Python 3**: The script is written in Python 3.
- **Markdown Files**: The script targets files with a `.md`
  extension.
- **YAML and H1 Detection**: Uses regular expressions to detect
  the presence of YAML frontmatter and an optional H1 heading.

## Customization

- **Directory Scanning**: By default, the script scans the
  current working directory. You can modify the
  `current_directory` variable or adapt the
  `find_and_delete_empty_files()` function to scan a different
  directory.
- **Criteria for Empty Files**: The script uses regex patterns
  to define what constitutes an empty file. You can adjust these
  patterns (`yaml_pattern` and `h1_pattern`) within the
  `is_empty_file()` function as needed.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.

# extract_info

# Energy Usage Data Extraction Script

## Overview

This Python script processes a collection of text files
containing energy usage data, extracting relevant details such
as estimated bills, total weekly kWh usage, and daily kWh usage.
The script then consolidates this data into a CSV file for easy
analysis and also logs detailed debug information for
verification.

## Features

- **Data Extraction**: Parses text files to extract:
  - Estimated bill amounts.
  - Total weekly energy usage in kWh.
  - Daily energy usage for each day of the week.
  - Start and end dates of the usage week.
- **CSV Output**: Aggregates the extracted data into a
  structured CSV file.
- **Debug Log**: Logs detailed extraction information to a debug
  file for troubleshooting.

## Usage

1. **Directory Setup**: Place all text files (`.txt`) containing
   energy usage data in a directory that the script will
   process. By default, the script processes files in the
   current working directory.

2. **Run the Script**:
   
   ```bash
   python3 energy_usage_extractor.py
   ```
   
   The script will:
   
   - Extract data from all `.txt` files in the directory.
   - Generate a CSV file (`collated_data.csv`) with the
     extracted data.
   - Create a debug log file (`debug_log.txt`) containing
     detailed information about the extracted content.

3. **CSV Output**: The resulting CSV file (`collated_data.csv`)
   will contain the following columns:
   
   - `week_start`: Start date of the usage week.
   - `week_end`: End date of the usage week.
   - `estimated_bill`: Estimated bill amount in dollars.
   - `total_kwh`: Total energy usage in kWh for the week.
   - `Sun`, `Mon`, `Tue`, `Wed`, `Thu`, `Fri`, `Sat`: Daily kWh
     usage for each day of the week.

## Requirements

- **Python 3**: The script is written in Python 3 and requires
  standard libraries (`os`, `re`, `csv`).
- **Text Files**: Ensure the input files are in `.txt` format
  and contain energy usage details in the expected structure.

## Customization

- **Directory Path**: Modify `directory_path` to specify a
  different directory for the input files:
  
  ```python
  directory_path = "/path/to/your/directory"
  ```

- **Output File Names**: Change `output_file` and `debug_file`
  to customize the names of the generated CSV and debug log
  files:
  
  ```python
  output_file = "your_output_file.csv"
  debug_file = "your_debug_file.txt"
  ```

## Debug Information

The script logs extracted data from each file in
`debug_log.txt`, which includes:

- Filename processed.
- Extracted details such as estimated bill, total kWh usage, and
  daily kWh usage.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.

# i3-dynamic-splits

# Dynamic i3 Window Splitter

## Overview

This Python script dynamically adjusts the window split
direction in the i3 window manager based on the dimensions of
the focused container. It uses the `i3ipc` library to interact
with the i3 tree structure, determining whether to split a
window horizontally or vertically for optimal layout.

## Features

- **Automatic Split Direction**: Determines the split direction
  based on the dimensions of the parent container:
  - **Horizontal Split** (`split h`): If the height of the
    container is greater than its width.
  - **Vertical Split** (`split v`): If the width of the
    container is greater than its height.
- **Tab/Stack Layout Navigation**: If the focused window is
  within a tabbed or stacked layout, the script navigates up to
  the appropriate parent container for determining the split
  direction.

## Usage

1. **Install the `i3ipc` Python Package**: This script requires
   the `i3ipc` library to interact with the i3 window manager.
   Install it via `pip`:
   
   ```bash
   pip install i3ipc
   ```

2. **Run the Script**: Execute the script to apply the dynamic
   split:
   
   ```bash
   python3 split_dynamic.py
   ```
   
   You can bind this script to a keyboard shortcut in your i3
   config file (`~/.config/i3/config`):
   
   ```plaintext
   bindsym $mod+s exec python3 /path/to/split_dynamic.py
   ```
   
   This allows you to trigger the script using a key combination
   (e.g., `$mod+s`).

## Requirements

- **Python 3**: The script is written in Python 3.
- **i3 Window Manager**: The script is designed to work with the
  i3 window manager, leveraging its tree layout for window
  management.
- **`i3ipc` Library**: Required to interact with i3's IPC
  (Inter-Process Communication) interface.

## Customization

- **Modify Split Logic**: You can adjust the logic that
  determines the split direction based on other conditions by
  editing the `split_dynamic` function.
- **Keyboard Shortcut**: Customize the key binding in your i3
  config to trigger this script as desired.

## Notes

- **Parent Container Check**: The script automatically navigates
  up to the parent container if the focused window is part of a
  tabbed or stacked layout.
- **Script Location**: Make sure to provide the correct path to
  the script in your i3 config file when setting up a keyboard
  shortcut.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.

# json-data

# Speed Test Data Analyzer

## Overview

This Python script loads speed test results from a JSON file,
analyzes the data to compute average download/upload speeds and
ping times, and visualizes the download speed trend over time
using a line plot.

## Features

- **Load Speed Test Data**: Reads speed test results from a JSON
  file.
- **Data Analysis**: Computes:
  - Average download speed (Mbps).
  - Average upload speed (Mbps).
  - Minimum and maximum ping times (ms).
- **Visualization**: Plots download speeds over time to provide
  a visual representation of speed trends.

## Usage

1. **Prepare Your JSON Data**: Ensure you have a JSON file named
   `speedtest_results.json` in the same directory as the script,
   or modify the `load_data` function to use a different file
   path. The JSON structure should be an array of entries
   containing `timestamp`, `download`, `upload`, and `ping`
   fields, for example:
   
   ```json
   [
       {
           "timestamp": "2024-10-01 10:00:00",
           "download": "50.5 Mbit/s",
           "upload": "10.5 Mbit/s",
           "ping": "20 ms"
       },
       ...
   ]
   ```

2. **Install Dependencies**: This script requires `matplotlib`
   for plotting:
   
   ```bash
   pip install matplotlib
   ```

3. **Run the Script**:
   
   ```bash
   python3 speedtest_analyzer.py
   ```
   
   The script will:
   
   - Load the JSON data.
   - Analyze the data and print the average download/upload
     speeds and min/max ping.
   - Display a plot of download speeds over time.

## Requirements

- **Python 3**: The script is written in Python 3.
- **`matplotlib`**: Used for plotting download speed trends.
  Install using `pip`.
- **JSON Data File**: Ensure the JSON data is correctly
  formatted and contains the required fields.

## Customization

- **Change Data File**: If your data is not in
  `speedtest_results.json`, modify the `load_data` function to
  point to the correct file:
  
  ```python
  data = load_data("/path/to/your/data.json")
  ```

- **Adjust Data Parsing**: The script expects download and
  upload speeds in "Mbit/s" and ping times in "ms". Adjust the
  parsing in the `analyze_data` function if your format differs.

## Notes

- **Error Handling**: The script logs errors for file handling,
  JSON decoding issues, and value parsing problems, making it
  easier to debug data inconsistencies.
- **Data Visualization**: The script currently only plots
  download speeds. You can extend the `plot_data` function to
  plot other metrics such as upload speed or ping times.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.

# renamer

# File and Directory Renaming Script

## Overview

This Python script recursively renames files and directories
within a given directory to adhere to certain naming
conventions. It aims to standardize file and folder names by
converting them to lowercase, replacing special characters, and
ensuring compatibility with Windows naming restrictions.

## Features

- **Naming Conventions**:
  - Converts all names to lowercase, except for `README.md`
    which is kept in uppercase.
  - Skips renaming of files and directories in CamelCase or
    PascalCase.
- **Special Character Handling**:
  - Replaces `&` with `and`.
  - Removes quotes, apostrophes, and certain punctuation
    (`.,!?()`) from names.
  - Replaces colons, hyphens, en dashes, and em dashes with
    underscores (`_`).
  - Replaces Windows-incompatible characters (`<>:"/\|?*`)
    with underscores.
  - Collapses multiple underscores into a single underscore
    and trims leading or trailing underscores.
- **Reserved Windows Names**: Appends an underscore (`_`) to any
  file or directory name that matches a reserved Windows name
  (e.g., `CON`, `PRN`, `AUX`, `COM1`).
- **Lowercase Extensions**: Converts all file extensions to
  lowercase.

## Usage

1. **Run the Script**: Use the script with a target directory as
   a command-line argument, or without arguments to process the
   current directory:
   
   ```bash
   python3 rename_files_and_directories.py <target_directory>
   ```
   
   - If `<target_directory>` is omitted, the script uses the
     current working directory.

2. **Process**: The script will:
   
   - Recursively traverse through all non-hidden files and
     directories.
   - Rename files and directories according to the specified
     conventions.
   - Skip any CamelCase or PascalCase names.

## Requirements

- **Python 3**: The script is written in Python 3.
- **File System Access**: The script needs read and write
  permissions to rename files and directories within the
  specified target directory.

## Customization

- **Excluding Certain Names**: The script currently preserves
  CamelCase and PascalCase names. If you want to rename these as
  well, modify the `is_camel_or_pascal_case` function.
- **Character Replacement Rules**: The character replacement
  logic can be customized in the `rename_item` function, where
  the script replaces characters like `&` and `<>:"/\|?*`.

## Notes

- **Handling Reserved Windows Names**: The script uses a list of
  reserved names (`CON`, `PRN`, `AUX`, `NUL`, etc.) to append an
  underscore if a name matches any reserved word.
- **Hidden Files and Directories**: The script ignores hidden
  files and directories (names starting with `.`).

## Example

Before running the script:

```rust
example &file!.txt
CamelCaseFile.md
CON.txt
some-folder/
```

After running the script:

```rust
example_andfile.txt
CamelCaseFile.md
CON_.txt
some_folder/
```

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.

# speaker

# Transcription Formatting Script

## Overview

This Python script formats a transcription from a Markdown file,
focusing on enhancing readability by organizing speaker turns,
timestamps, and text content. It processes the transcription
part of the file, adds structure for speaker changes, and
outputs a formatted version of the transcription.

## Features

- **Speaker Turn Marking**: Splits the transcription based on
  `[SPEAKER_TURN]` markers to distinguish different speakers.
- **Timestamp Extraction**: Finds and organizes timestamps
  within the transcription.
- **Content Formatting**: Formats text for each timestamp by
  ensuring proper spacing and structure, removing unnecessary
  backslashes, and adding line breaks for clarity.

## Usage

1. **Run the Script**: Execute the script by passing the path to
   a Markdown file containing the transcription:
   
   ```bash
   ./speaker.py <input_markdown_file>
   ```
   
   Replace `<input_markdown_file>` with the path to your
   Markdown file.
   
   The script will create a new formatted Markdown file with
   `_formatted` appended to its name, e.g., `transcription.md`
   will become `transcription_formatted.md`.

2. **Output**: The formatted file will have:
   
   - Each `[SPEAKER_TURN]` on its own line.
   - Timestamps and corresponding text content organized for
     readability.
   - Removed unnecessary backslashes and added appropriate line
     breaks.

## Requirements

- **Python 3**: The script is written in Python 3.
- **Input File Format**: The script expects a Markdown file with
  a transcription section containing `[SPEAKER_TURN]` markers,
  timestamps (in brackets), and text content.

## Customization

- **Regex Modifications**: If your transcription format varies
  or requires different processing rules, adjust the
  `format_transcription()` function, particularly the regex
  patterns used to split the text and extract timestamps.

## Debugging Information

The script includes debugging print statements to help you
understand how parts of the transcription are processed:

- **Processing Part**: Displays each segment being processed.
- **Matched Timestamps and Content**: Shows how timestamps and
  content are extracted and formatted.
- **Input and Output Debugging**: Prints both the original and
  formatted text for comparison.

To disable debugging output, simply remove or comment out the
`print()` statements in the `format_transcription()` and
`main()` functions.

## Example

**Input File**: `transcription.md`

```rust
---
title: Example Transcription
date: 2024-10-01
---

[SPEAKER_TURN]
[00:00:01] Hello, everyone.
[00:00:03] Welcome to the show.

[SPEAKER_TURN]
[00:00:10] Thank you for having me.
```

**Output File**: `transcription_formatted.md`

```rust
[SPEAKER_TURN]
[00:00:01]
Hello, everyone.

[00:00:03]
Welcome to the show.

[SPEAKER_TURN]
[00:00:10]
Thank you for having me.
```

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.

# speedtest

# Speedtest Logging Script

## Overview

This Python script runs a network speed test using
`speedtest-cli` and logs the results (ping, download, and upload
speeds) in a JSON file. Each time the script is executed, it
appends a new entry with the current timestamp and speed test
results to the file.

## Features

- **Automated Speed Test**: Runs a speed test using
  `speedtest-cli` and extracts `ping`, `download`, and `upload`
  results.
- **JSON Logging**: Appends the speed test results to a JSON
  file in a specified directory, maintaining a history of
  network performance over time.
- **Timestamped Entries**: Adds a timestamp to each entry for
  easy tracking of results over time.

## Usage

1. **Install `speedtest-cli`**: Make sure `speedtest-cli` is
   installed on your system. You can install it using `pip`:
   
   ```bash
   pip install speedtest-cli
   ```

2. **Run the Script**: Execute the script manually or schedule
   it to run periodically (e.g., using `cron` on Linux):
   
   ```bash
   python3 speedtest_logger.py
   ```

3. **Log Directory and File**:
   
   - The results are stored in
     `/home/jwl/Documents/speedtests/speedtest_results.json`.
     If the directory does not exist, it will be created
     automatically.
   - You can change the directory and file name by modifying
     the `directory` and `filename` variables at the top of the
     script.

## JSON File Structure

The JSON file contains an array of entries, each with:

- **timestamp**: The date and time of the speed test.
- **ping**: The latency in milliseconds (e.g., `"15 ms"`).
- **download**: The download speed in megabits per second (e.g.,
  `"50 Mbit/s"`).
- **upload**: The upload speed in megabits per second (e.g.,
  `"10 Mbit/s"`).

Example:

```json
[
    {
        "timestamp": "2024-10-01 12:00:00",
        "ping": "15 ms",
        "download": "50 Mbit/s",
        "upload": "10 Mbit/s"
    },
    {
        "timestamp": "2024-10-02 12:00:00",
        "ping": "20 ms",
        "download": "45 Mbit/s",
        "upload": "12 Mbit/s"
    }
]
```

## Customization

- **Change Output Directory or File**: Modify the `directory`
  and `filename` variables at the start of the script:
  
  ```python
  directory = "/your/desired/directory"
  filename = "your_results.json"
  ```

- **Error Handling**: The script checks for a JSON decoding
  error when loading the existing data. If the file is not a
  valid JSON, it starts with an empty array.

## Automation

To run the script automatically at regular intervals, you can
use a scheduling tool like `cron`. For example, to run the
script every hour, add the following line to your crontab:

```bash
0 * * * * /usr/bin/python3 /path/to/speedtest_logger.py
```

## Requirements

- **Python 3**: The script is written in Python 3.
- **`speedtest-cli`**: Required to run the speed tests. Install
  via `pip install speedtest-cli`.
- **File System Access**: Ensure the script has write
  permissions to the specified directory.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.

# status_bar

# Directory Scanning Script with Progress Bar

## Overview

This Python script scans through a specified directory and its
subdirectories to collect metadata (file paths and modification
times) for all non-hidden files. It uses the `tqdm` library to
display a progress bar, providing a visual indication of the
scan's progress.

## Features

- **Directory Traversal**: Recursively scans the specified
  directory and its subdirectories.
- **File Metadata Collection**: Collects file paths and
  modification times (`mtime`).
- **Progress Bar**: Uses `tqdm` to display a progress bar during
  the scan, showing how many files have been processed.

## Usage

1. **Install `tqdm`**: Ensure you have the `tqdm` library
   installed, which is used for the progress bar. You can
   install it with:
   
   ```bash
   pip install tqdm
   ```

2. **Run the Script**: Import the function
   `scan_directory_with_progress()` in your script, or call it
   directly within a `main()` function:
   
   ```python
   from scanner import scan_directory_with_progress
   
   directory = "/path/to/your/directory"
   files_metadata = scan_directory_with_progress(directory)
   ```
   
   Replace `"/path/to/your/directory"` with the path of the
   directory you wish to scan.

3. **Output**: The function returns a list of tuples containing:
   
   - **File Path**: The full path to each file.
   - **Modification Time**: The last modified time of each file
     (Unix timestamp).

## Example

```python
if __name__ == "__main__":
    directory = "/home/user/my_directory"
    metadata = scan_directory_with_progress(directory)

    for file_path, mtime in metadata:
        print(f"File: {file_path}, Last Modified: {mtime}")
```

This will print each file's path and its modification time,
while showing a progress bar as the directory is being scanned.

## Requirements

- **Python 3**: The script is written in Python 3.

- **`tqdm` Library**: For displaying the progress bar. Install
  it using:
  
  ```bash
  pip install tqdm
  ```

## Customization

- **Hidden File Exclusion**: The script ignores hidden files
  (those starting with `.`). You can modify the `if not
  file.startswith('.')` condition in the loop to include hidden
  files if desired.
- **Additional Metadata**: You can extend the `files_metadata`
  collection to include other file attributes like size (
  `os.path.getsize(file_path)`) or permissions
  (`os.stat(file_path).st_mode`).

## Notes

- **Progress Bar**: The progress bar will display the total
  number of files being processed. If the directory contains a
  large number of files, the progress bar provides real-time
  feedback on the script's progress.
- **Cross-Platform**: The script uses `os.walk` and `os.path`
  for file traversal and metadata collection, making it
  compatible with various operating systems (Linux, macOS,
  Windows).

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.

# batch_transcribe

# Audio Conversion and Transcription Script

## Overview

This Bash script automates the process of converting audio files
to 16 kHz WAV format and then transcribing them using a speech
recognition model. It uses `ffmpeg` for audio conversion and
`whisper.cpp` for transcription. The transcriptions are saved as
text files for easy access.

## Features

- **Audio Conversion**: Converts all audio files in the
  specified directory to 16 kHz WAV format, stored temporarily.
- **Transcription**: Transcribes the converted WAV files using a
  specified model and saves the output in text format.
- **Progress Tracking**: Displays progress updates during both
  conversion and transcription steps.
- **Clean-up**: Automatically removes the temporary directory
  used for converted files after processing.

## Usage

1. **Set up directories and paths**:
   
   - **AUDIO_DIR**: Directory containing your original audio
     files (`./audio` by default).
   - **TEMP_DIR**: Temporary directory to store converted WAV
     files (`./converted/` by default).
   - **OUTPUT_DIR**: Directory to store transcription output
     (`./transcriptions/` by default).
   - **MODEL_PATH**: Path to the speech recognition model file
     (`ggml-small.en-tdrz.bin`).
   - **MAIN_EXEC**: Path to the main executable for
     transcription (`whisper.cpp`).

2. **Run the script**:
   
   ```bash
   ./your_script_name.sh
   ```

3. The script will:
   
   - Convert all audio files in `AUDIO_DIR` to WAV format.
   - Transcribe each WAV file using `whisper.cpp` with the
     `--tinydiarize` option.
   - Save the transcriptions as `.txt` files in `OUTPUT_DIR`.
   - Clean up the temporary converted files after processing.

## Requirements

- **Bash**: The script is designed to be run in a Bash shell.
- **ffmpeg**: Required for converting audio files to the desired
  format (16 kHz WAV).
- **whisper.cpp**: A speech recognition tool for transcribing
  audio. Ensure the paths to the model and main executable are
  correct.
- **Audio Files**: The script processes all files in the
  `AUDIO_DIR` directory, so ensure it contains supported audio
  formats (e.g., `.mp3`, `.wav`, etc.).

## Customization

- **Audio Directory**: Modify the `AUDIO_DIR` variable to point
  to your desired directory containing audio files.
- **Model and Executable Path**: Adjust `MODEL_PATH` and
  `MAIN_EXEC` to match your environment's setup.
- **Output Locations**: The output transcriptions will be saved
  in `OUTPUT_DIR`, which can be changed as needed.

## Example

```bash
./audio-to-text.sh
```

This will convert all audio files in `./audio/`, transcribe
them, and save the transcriptions to `./transcriptions/`.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.

# generate_readme

# Index Generation Script

## Overview

This Bash script automates the creation and updating of
`README.md` files in each directory of a vault or repository,
except the root directory. It provides a file count, lists
categories, generates a Table of Contents (TOC) for
subdirectories, and organizes notes or files for better
navigation within the vault.

## Features

- **Creates `README.md` Files**: Generates a `README.md` file in
  each directory if it doesn't exist.
- **Updates `README.md` Files**: Populates each `README.md`
  with:
  - A unique hash as a comment (for tracking changes).
  - A header with the directory name.
  - A count of the files in the directory.
  - A Table of Contents for subdirectories.
  - A list of files in each directory and its subdirectories.
- **Excludes Hidden Files/Directories**: Ignores files or
  directories that are hidden (names starting with `.`).

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
   
   - **`README.md` Files Created**: A new `README.md` will be
     created in directories that don't have one.
   - **`README.md` Files Updated**: Existing `README.md` files
     will be updated with a list of notes/files and a TOC for
     any subdirectories.

## Requirements

- **Bash**: The script is intended for a Bash shell.
- **`realpath` Utility**: Used to resolve and normalize
  directory paths.
- **`find`, `sed`, `iconv`, `md5sum` Utilities**: Standard
  command-line tools used for directory traversal, string
  manipulation, and hashing.

## Customization

- **Modify `VAULT_PATH`**: The base directory for processing is
  set to the current working directory. You can modify this by
  changing the `VAULT_PATH` variable:
  
  ```bash
  VAULT_PATH="/your/desired/path"
  ```

- **File Naming and TOC Generation**: You can adjust the TOC
  generation and the file naming convention by tweaking the
  `generate_readme` and `to_snake_case` functions.

## Notes

- **`README.md` Structure**: Each `README.md` file contains:
  - A header with the directory name and a count of
    notes/files.
  - A list of all files in the directory, excluding hidden
    files.
  - A categorized TOC listing any subdirectories and their
    notes.
- **Skipping Root Directory**: The script does not generate a
  `README.md` for the root directory.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.

# info

# Directory Information Script

## Overview

This Bash script provides an overview of the contents of the
current directory. It reports the total size of the directory,
the number of files and folders, and identifies any empty files
or directories.

## Features

- **Current Directory Information**: Uses the current directory
  as the target for analysis.
- **Total Size Calculation**: Displays the total size of the
  directory in gigabytes.
- **File and Folder Count**: Counts and displays the total
  number of files and folders.
- **Empty File and Directory Identification**: Provides a count
  of any empty files or directories within the current
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

- **Directory Information for**: Displays the current directory
  being analyzed.
- **Total Size (in GB)**: The total size of the directory and
  its contents in gigabytes.
- **Total Files**: The count of all files (including those in
  subdirectories).
- **Total Folders**: The count of all folders (including
  subdirectories).
- **Total Empty Files**: The count of files that are empty (0
  bytes in size).
- **Total Empty Directories**: The count of directories that do
  not contain any files or subdirectories.

## Requirements

- **Bash**: The script should be run in a Bash shell.
- **`du`, `find`, `awk`, and `wc` utilities**: These standard
  command-line tools are used for directory size calculations,
  file/folder counting, and filtering.

## Customization

You can modify the `DIR` variable to point to any specific
directory if you don't want to use the current directory:

```bash
DIR="/path/to/your/directory"
```

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.

# symlinks

# Emacs Configuration Symlink Script

## Overview

This Bash script creates symbolic links (symlinks) from an Emacs
configuration directory within a dotfiles repository to the
user's home `.emacs.d` directory. This approach makes it easy to
manage and version control Emacs configuration files.

## Features

- **Automatic Directory Creation**: Ensures that the target
  `.emacs.d` directory exists in the user's home directory.
- **Symlink Creation**: Creates symbolic links for specific
  files or directories from the source dotfiles directory to the
  target Emacs configuration directory.
- **Easily Manageable Configuration**: Allows for easy
  management of Emacs configuration files through symlinks,
  making updates and backups more efficient.

## Usage

1. **Define the Files to be Symlinked**:
   
   - Modify the `FILES` array in the script to include the
     names of the files and directories you want to symlink:
     
     ```bash
     FILES=("init.el" "lisp" "config.org")
     ```
     
     Each entry should match the relative path of a file or
     directory inside your `dotfiles/.emacs.d` directory.

2. **Run the Script**:
   
   - Make sure the script is executable:
     
     ```bash
     chmod +x symlink_emacs.sh
     ```
   
   - Execute the script to create the symlinks:
     
     ```bash
     ./symlink_emacs.sh
     ```
   
   The script will create symlinks for each file/directory in
   the `FILES` array from `~/dotfiles/.emacs.d` to `~/.emacs.d`.

## Example

**Directory Structure**:

- Source directory (`~/dotfiles/.emacs.d/`):
  
  ```rust
  init.el
  lisp/
  config.org
  ```

- `FILES` array:
  
  ```bash
  FILES=("init.el" "lisp" "config.org")
  ```

After running the script, your `~/.emacs.d/` directory will
contain symlinks pointing to the corresponding files and
directories in `~/dotfiles/.emacs.d/`.

## Requirements

- **Bash Shell**: The script is designed to be run in a Bash
  shell.
- **Read and Write Permissions**: Ensure that you have the
  necessary permissions to create directories and symlinks in
  the target location.

## Customization

- **Source and Target Directories**: Modify `SOURCE_DIR` and
  `TARGET_DIR` if your dotfiles or Emacs configuration are
  located in different directories:
  
  ```bash
  SOURCE_DIR="/path/to/your/dotfiles/.emacs.d"
  TARGET_DIR="/path/to/your/.emacs.d"
  ```

- **Symlink Files and Directories**: Adjust the `FILES` array to
  include all files and directories you wish to symlink.

## Notes

- **Avoid Overwriting Existing Files**: If the target files
  already exist, the script may fail to create the symlinks.
  Make sure to backup or remove existing files if necessary.
- **Safe to Run Multiple Times**: If new files are added to
  `FILES`, rerun the script to create symlinks for the new
  additions.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.

# til

# Directory-Named README Generator Script

## Overview

This Bash script generates or updates `README.md` files for each
directory in a specified vault or repository path. Each
`README.md` file summarizes the contents of its corresponding
directory, providing an overview of all files and
subdirectories. The script uses "kebab case" for naming and
formatting and creates Markdown-friendly links.

## Features

- **Automatic README Creation**: Creates a `README.md` file in
  each directory, named after the directory itself in kebab
  case.
- **Directory and File Listing**: Lists all files and
  subdirectories in the `README.md`, providing a count of items
  in each subdirectory and linking to each file.
- **URL Encoding for Links**: Converts spaces to `%20` in URLs
  for compatibility with GitHub Markdown.
- **Kebab Case Naming**: Converts file and directory names to
  kebab case to maintain consistent naming conventions.

## Usage

1. **Set the Vault/Repository Path**: Ensure the `VAULT_PATH`
   variable in the script points to the base directory where you
   want the `README.md` files to be created/updated:
   
   ```bash
   VAULT_PATH="/path/to/your/vault"
   ```

2. **Run the Script**: Make the script executable and then
   execute it:
   
   ```bash
   chmod +x update_readmes.sh
   ./update_readmes.sh
   ```
   
   The script will traverse all non-hidden directories within
   `VAULT_PATH`, creating or updating a `README.md` file in each
   one.

3. **Output**: Each `README.md` will:
   
   - Have a header titled "Overview of `<directory-name>`".
   - List all subdirectories with a count of items within them.
   - Include kebab case links to each file and subdirectory,
     with appropriate formatting for Markdown.

## Example

**Directory Structure**:

```rust
/projects/
    /my-first-project/
        script.sh
        README.md
    /another_project/
        data.txt
        config.yml
```

After running the script, `README.md` files will be generated
for each directory:

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
- **Read and Write Permissions**: Ensure the script has the
  necessary permissions to read from the source directory and
  write to each target directory.
- **Standard Utilities**: The script uses common Linux utilities
  like `realpath`, `sed`, `find`, and `iconv`.

## Customization

- **Exclude Files/Directories**: By default, hidden files (names
  starting with `.`) and `README.md` files are excluded. You can
  customize this behavior by modifying the conditions in the
  `update_readme()` function.
- **Kebab Case Conversion**: If you wish to use a different
  naming convention, modify the `to_kebab_case()` function as
  necessary.

## Notes

- **Directory Naming in README.md**: Each `README.md` is named
  after its parent directory in kebab case.
- **Non-Markdown Files**: If a file is not a Markdown file
  (`.md`), its extension is included in the display name;
  otherwise, the extension is removed.
- **Script Recursion**: The script recursively traverses all
  directories from `VAULT_PATH` but skips hidden directories and
  files.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.

# watch_cronjobs

# Update Crontab Script

## Overview

This Bash script updates the user's crontab using a predefined
crontab file. It allows for easy management of cron jobs by
maintaining them in a version-controlled file, which can be
applied to the system's crontab with a single command.

## Features

- **Automated Crontab Update**: Automatically sets the user's
  crontab to match the contents of a specified crontab file.
- **File Existence Check**: Verifies that the crontab file
  exists before attempting to update the crontab, providing an
  error message if the file is not found.
- **Centralized Crontab Management**: Makes it easy to version
  control and update cron jobs by storing them in a single file.

## Usage

1. **Set the Crontab File Path**: Make sure the `CRONTAB_FILE`
   variable in the script points to the correct path of your
   crontab file:
   
   ```bash
   CRONTAB_FILE="/path/to/your/cronjobs"
   ```

2. **Add or Modify Cron Jobs**: Edit the crontab file (e.g.,
   `/home/jwl/dotfiles/.scripts/cronjobs`) to add or modify cron
   jobs as needed. The file should follow standard crontab
   syntax.
   
   **Example crontab file**:
   
   ```rust
   0 5 * * * /home/jwl/scripts/backup.sh
   */15 * * * * /home/jwl/scripts/check_updates.sh
   ```

3. **Run the Script**: Make sure the script is executable and
   then run it to update your crontab:
   
   ```bash
   chmod +x update_crontab.sh
   ./update_crontab.sh
   ```
   
   If the crontab file is found, the script will apply the cron
   jobs defined in that file to the user's crontab.

4. **Output**:
   
   - If the crontab file is found, the crontab will be updated
     without any further message.
   
   - If the crontab file is missing, the script will print an
     error message:
     
     ```rust
     Crontab file does not exist: /home/jwl/dotfiles/.scripts/cronjobs
     ```

## Requirements

- **Bash Shell**: The script is intended to run in a Bash shell.
- **Crontab Access**: Ensure you have permissions to update the
  user's crontab.
- **Valid Crontab File**: The crontab file must be in a valid
  format for cron to accept it.

## Customization

- **Change the Crontab File Path**: If your crontab file is
  located elsewhere, modify the `CRONTAB_FILE` variable:
  
  ```bash
  CRONTAB_FILE="/path/to/your/custom_cron_file"
  ```

## Notes

- **Error Handling**: The script checks for the existence of the
  crontab file before attempting to update the crontab. If the
  file does not exist, it will exit with an error.
- **Version Control**: By keeping your cron jobs in a file
  (e.g., within a Git repository), you can easily track changes
  and revert to previous versions if necessary.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.

# getter

# Script Index Generator

## Overview

This Python script scans a directory for various script files
(such as `.sh`, `.py`, `.go`, `.rs`, `.pl`, `.rb`) and generates
a `README.md` file containing basic information about each
script. The README file includes the file path, shebang, and
extracted single-line comments to provide a description for each
script.

## Features

- **File Type Support**: Automatically detects and processes
  files with extensions `.sh`, `.py`, `.go`, `.rs`, `.pl`, and
  `.rb`.
- **Shebang Detection**: Extracts and displays the shebang
  (`#!`) from each script to indicate its interpreter.
- **Description Extraction**: Collects single-line comments from
  scripts to provide a brief description or notes about the
  functionality.
- **Non-Recursive**: Scans only the root directory without
  traversing into subdirectories.

## Usage

1. **Place the Script in the Target Directory**: Save this
   script in the directory containing the scripts you want to
   index.

2. **Run the Script**:
   
   ```bash
   python3 generate_script_readme.py
   ```
   
   The script will:
   
   - Scan the current directory for script files.
   - Extract metadata such as the shebang and comments from
     each file.
   - Generate a `README.md` file containing an index of all
     detected scripts.

3. **Output**:
   
   - A `README.md` file will be created in the same directory,
     with sections for each script that include:
     - **File Name**: The script's name.
     - **Path**: The path to the script file.
     - **Shebang**: The shebang line to indicate the
       interpreter.
     - **Description**: Comments extracted from the script,
       formatted for easy reading.

## Requirements

- **Python 3**: The script is written in Python 3.
- **Script Files**: The script targets files ending in `.sh`,
  `.py`, `.go`, `.rs`, `.pl`, and `.rb`. You can modify the
  extensions in the code to support other file types.

## Customization

- **Target Directory**: By default, the script scans the current
  directory (`"."`). To change this, modify the
  `generate_readme` function to point to a specific path:
  
  ```python
  generate_readme("/path/to/your/scripts")
  ```

- **File Extensions**: To add or remove file types, edit the
  list of supported extensions in the `generate_readme`
  function:
  
  ```python
  if os.path.isfile(file_path) and file.endswith((".sh", ".py", ".your_extension")):
  ```

## Notes

- **Single-Line Comments Only**: The script only extracts
  single-line comments starting with `#` or `//`.
- **Inline Code Formatting**: Comments are wrapped in backticks
  to appear as inline code in the generated README.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.

# doc_maker

# README File Generator Script

## Overview

This Python script scans the current directory and creates an
empty Markdown README file (`.md`) for each file that does not
already have a corresponding README. It ensures that each file
in the directory has a README file with the same base name.

## Features

- **Automatic README Creation**: For each file in the directory,
  an empty Markdown file (`.md`) is created if it doesn't
  already exist.
- **File Matching by Name**: Each README is named after the
  corresponding file, with the same base name but with a `.md`
  extension.
- **Skip Existing README Files**: The script checks for existing
  README files and skips their creation if they are already
  present.

## Usage

1. **Run the Script**: Navigate to the desired directory and
   execute the script:
   
   ```bash
   python3 create_readmes.py
   ```
   
   This will generate `.md` files for each file without an
   existing README in the current directory.

2. **Output**:
   
   - **Created**: If a README file is created, you will see a
     message like:
     
     ```rust
     Created: <filename>.md
     ```
   
   - **Skipped**: If a README file already exists, you will see
     a message like:
     
     ```rust
     Skipped: <filename>.md already exists
     ```

## Requirements

- **Python 3**: The script is written in Python 3.
- **Directory Scanning**: The script scans the current working
  directory for all files and creates README files accordingly.

## Customization

- **Target Directory**: By default, the script works in the
  current directory. To change this behavior, modify
  `current_directory` to point to a specific path:
  
  ```python
  current_directory = "/path/to/your/directory"
  ```

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.

# cronjobs

# Cron Jobs Overview

## Overview

This README provides a summary of the cron jobs configured for
automation tasks on your system. These jobs perform regular file
backups, speed tests, and data fetching, each running at a
specified time of day.

## Cron Job Schedule

1. **Backup `.zsh_history` File**
   
   - **Time**: Every hour at the 5th minute.
   
   - **Command**:
     
     ```bash
     cp /home/jwl/.zsh_history /home/jwl/dotfiles/
     ```
   
   - **Description**: Copies the `.zsh_history` file to the
     `dotfiles` directory for backup and versioning.

2. **Run Speed Test Script**
   
   - **Time**: Every hour at the 6th minute.
   
   - **Command**:
     
     ```bash
     /home/jwl/.scripts/speedtest.py
     ```
   
   - **Description**: Executes a Python script to perform a
     speed test. The script is located in `.scripts` directory.

3. **Log Speed Test Results**
   
   - **Time**: Every hour at the 35th minute.
   
   - **Command**:
     
     ```bash
     /usr/bin/python3 /home/jwl/projects/speedtests/python/speedtest.py >> /home/jwl/projects/speedtests/logs/speedtest.log 2>&1
     ```
   
   - **Description**: Runs the Python speed test script and
     appends the output (including errors) to `speedtest.log`.

4. **Fetch Speed Test Data**
   
   - **Time**: Every hour at the start of the hour.
   
   - **Command**:
     
     ```bash
     /home/jwl/projects/speedtests/frontend/speedtest-results/fetch-data.sh >> /home/jwl/projects/speedtests/logs/fetch-data.log 2>&1
     ```
   
   - **Description**: Runs a shell script to fetch speed test
     data and appends the output (including errors) to
     `fetch-data.log`.

## Crontab Format Reference

```rust
<Minute> <Hour> <Day of Month> <Month> <Day of Week> <Command>
```

For all the cron jobs above:

- **Minute**: The minute of each hour when the job is run.
- **Hour**: The hour of the day (24-hour format) when the job is
  run.
- **Command**: The specific command to execute, including any
  redirection of output.

## Notes

- **Logging**: Two of the jobs redirect their output and errors
  to log files in the `logs` directory.
- **Script Paths**: Make sure all script paths (`speedtest.py`,
  `fetch-data.sh`) are correct and the scripts have execute
  permissions.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.

# clone

# Git Repository Cloning Script

## Overview

This Bash script clones all Git repositories in the current
directory to a specified target directory. It scans for all
`.git` directories, extracts their parent folder names as
project names, and clones each repository to the target
location.

## Features

- **Bulk Repository Cloning**: Automatically identifies and
  clones all Git repositories in the current directory.
- **Custom Target Directory**: Allows the user to specify a
  target directory for cloned repositories.
- **Automatic Directory Creation**: Creates the target directory
  if it doesn't already exist.

## Usage

1. **Run the Script**: Use the following command to execute the
   script, replacing `<target-directory>` with your desired
   directory:
   
   ```bash
   ./clone_repos.sh <target-directory>
   ```
   
   - `<target-directory>`: The path where you want all
     repositories to be cloned.

2. The script will loop over all `.git` directories in the
   current directory, clone each repository, and store them in
   `<target-directory>`.

## Example

```bash
./clone_repos.sh ~/my_repos
```

This will clone all Git repositories from the current directory
to `~/my_repos`.

## Requirements

- **Bash**: The script should be run in a Bash shell.
- **Git**: Ensure that `git` is installed and accessible in your
  `PATH`.

## Customization

- **Target Directory**: Modify the script to change the default
  target directory or add additional logic for more complex
  directory structures.
- **Repository Detection**: The script assumes all
  subdirectories with `.git` are Git repositories. If you have a
  different structure, you may need to adjust the repository
  detection logic.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.

# README

# Directory Organizer Script

## Overview

This script scans a directory, organizes files by their base
names, and moves them into corresponding folders. If a file
shares the same base name (ignoring extensions) with other
files, all such files are moved into a newly created folder
named after their base name. The script skips all hidden files
and folders by default (those starting with a `.`).

For example, if the directory contains `example.py`,
`example.md`, and `example.sh`, the script will create a folder
named `example` and move all three files into this folder.

## Usage

1. **Ensure the script is executable**: You may need to change
   the permissions to make the script executable:
   
   ```bash
   chmod +x your_script_name.py
   ```

## License

These scripts are open-source and available for free use and modification under [The Unlicense](https://unlicense.org/), making them public domain and free to use without any restrictions.

# word_check

# Word Check Script

## Overview

These scripts, one written in Perl and the other in Python,
process a text file containing words, filter them based on
specific criteria, and output a cleaned list of words. The main
purpose of the scripts is to remove duplicates, words containing
certain patterns, and words that don't meet specific length or
character requirements.

## Features

- **Converts Words to Lowercase**: All words in the input file
  are converted to lowercase for uniformity.
- **Removes Punctuation**: Any punctuation within the words is
  stripped before processing.
- **Filters Words Based on Criteria**:
  - Must contain only English letters (A-Z, case-insensitive).
  - Must be at least 4 characters long.
  - Must not have three or more consecutive identical letters
    (e.g., "aaa").
  - Must not start or end with two or more identical letters
    (e.g., "aab", "baa").
- **Ensures Unique Words**: No duplicate words are included in
  the output.

## Input and Output

- **Input**: Both scripts read from an input file named
  `word_list.txt`.
- **Output**: The filtered words are saved to:
  - **Perl Script**: `clean_word_list.txt`
  - **Python Script**: `clean_list.txt`

## How to Use

### Perl Script

1. Place your word list in a file named `word_list.txt` in the
   same directory as the script.

2. Run the script:
   
   ```bash
   perl word_check.pl
   ```

3. The filtered and cleaned word list will be written to
   `clean_word_list.txt`.

### Python Script

1. Place your word list in a file named `word_list.txt` in the
   same directory as the script.

2. Run the script:
   
   ```bash
   python3 clean_word_list.py
   ```

3. The filtered and cleaned word list will be written to
   `clean_list.txt`.

## Requirements

### For the Perl Script

- **Perl**: Ensure Perl is installed on your system.
- **UTF-8 Encoding**: The input file should be UTF-8 encoded for
  proper processing.

### For the Python Script

- **Python 3**: The script is written for Python 3 and uses
  standard libraries (`string`, `re`).

## Customization

To change the input or output file names, modify the relevant
lines in either script:

- **Perl**:
  
  ```perl
  my $input_file = 'word_list.txt';
  my $output_file = 'clean_word_list.txt';
  ```

- **Python**:
  
  ```python
  input_file = "word_list.txt"
  output_file = "clean_list.txt"
  ```

Set `input_file` and `output_file` to any file paths of your
choice.

## Differences Between Perl and Python Versions

While both scripts achieve the same purpose, there are some
differences in implementation:

- **Function Organization**: The Python script uses separate
  functions to modularize checks for word validity, whereas the
  Perl script handles all checks within a single loop.
- **Output File Name**: The output file names are different by
  default (`clean_word_list.txt` for Perl and `clean_list.txt`
  for Python).

## License

These scripts are open-source and available for free use and modification under [The Unlicense](https://unlicense.org/),
making them public domain and free to use without any restrictions.
