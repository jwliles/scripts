Thank you for the details! Here's a revised README for your scripts directory, including all the scripts listed in your structure and with GitHub-compatible anchor links for easy navigation.

---

# Scripts Directory README

Welcome to the `scripts` directory! This README provides an overview of each script, its functionality, and usage. Use the table of contents below to quickly navigate through the different scripts and their descriptions.

## Table of Contents

- [Batch Transcribe](#batch-transcribe)
- [Change Detector](#change-detector)
- [Clone Repositories](#clone-repositories)
- [Cronjobs](#cronjobs)
- [Delete Empty Files](#delete-empty-files)
- [Documentation Maker](#documentation-maker)
- [Extract Info](#extract-info)
- [Generate README](#generate-readme)
- [File Getter](#file-getter)
- [i3 Dynamic Splits](#i3-dynamic-splits)
- [Info Script](#info-script)
- [JSON Data Analyzer](#json-data-analyzer)
- [File Mover](#file-mover)
- [Renamer Script](#renamer-script)
- [Speaker Formatter](#speaker-formatter)
- [Speedtest Logger](#speedtest-logger)
- [Status Bar Data](#status-bar-data)
- [Symlinks](#symlinks)
- [Today I Learned (TIL)](#today-i-learned-til)
- [Virtual Environment](#virtual-environment)
- [Watch Cronjobs](#watch-cronjobs)
- [Word Check](#word-check)

---

## Batch Transcribe

- **Files**: `batch_transcribe.sh`, `batch_transcribe.md`
- **Description**: A shell script to transcribe multiple audio files using a specified model and tool.
- **Usage**: Modify the script to set the audio directory, temporary directory for converted files, and the path to the transcription model and executable.
- **Run**: 
   ```bash
   ./batch_transcribe.sh
   ```
- **Output**: Creates transcriptions for each audio file in the specified output directory.

---

## Change Detector

- **Files**: `change_detector.py`, `change_detector.md`
- **Description**: A Python script that scans a directory to detect changes in file contents, generating and storing hashes.
- **Usage**: Modify the script to specify the directory to scan.
- **Run**: 
   ```bash
   python3 change_detector.py
   ```

---

## Clone Repositories

- **Files**: `clone.sh`, `clone.md`
- **Description**: Clones Git repositories found in the current directory to a target directory.
- **Usage**: Provide the target directory as an argument to the script.
- **Run**: 
   ```bash
   ./clone.sh <target-directory>
   ```

---

## Cronjobs

- **Files**: `cronjobs.txt`, `cronjobs.md`
- **Description**: A text file containing predefined cron jobs that can be loaded to the user's crontab.
- **Usage**: Manually add the jobs to your crontab or use a script to do so.

---

## Delete Empty Files

- **Files**: `delete-empty-files.py`, `delete-empty-files.md`
- **Description**: Deletes Markdown files containing only YAML frontmatter or a single H1 heading.
- **Usage**: Run the script in the target directory.
- **Run**: 
   ```bash
   python3 delete-empty-files.py
   ```

---

## Documentation Maker

- **Files**: `doc_maker.py`, `doc_maker.md`
- **Description**: Generates an index of all scripts in the current directory with descriptions extracted from their comments.
- **Usage**: Run the script in the target directory.
- **Run**: 
   ```bash
   python3 doc_maker.py
   ```

---

## Extract Info

- **Files**: `extract_info.py`, `extract_info.md`
- **Description**: Extracts specific information such as bill estimates and energy usage from text files.
- **Usage**: Modify the input directory and run the script to parse and save the extracted data.
- **Run**: 
   ```bash
   python3 extract_info.py
   ```

---

## Generate README

- **Files**: `generate_readme.sh`, `generate_readme.md`
- **Description**: Creates a `README.md` file for each directory, summarizing its contents.
- **Usage**: Set the `VAULT_PATH` and run the script to generate the `README.md` files.
- **Run**: 
   ```bash
   ./generate_readme.sh
   ```

---

## File Getter

- **Files**: `getter.py`, `getter.md`
- **Description**: Generates empty `README.md` files for each script in the current directory if they do not already exist.
- **Usage**: Run the script in the target directory.
- **Run**: 
   ```bash
   python3 getter.py
   ```

---

## i3 Dynamic Splits

- **Files**: `i3-dynamic-splits.py`, `i3-dynamic-splits.md`
- **Description**: A Python script to dynamically split windows based on dimensions in the i3 window manager.
- **Usage**: Requires the `i3ipc` Python library and the i3 window manager.
- **Run**: 
   ```bash
   python3 i3-dynamic-splits.py
   ```

---

## Info Script

- **Files**: `info.sh`, `info.md`
- **Description**: A shell script to display information about a specified directory, such as total size, number of files, and empty directories.
- **Usage**: Run the script to display the summary.
- **Run**: 
   ```bash
   ./info.sh
   ```

---

## JSON Data Analyzer

- **Files**: `json-data.py`, `json-data.md`
- **Description**: A Python script to analyze JSON data and plot the results, typically for network speed tests.
- **Usage**: Requires the `matplotlib` library.
- **Run**: 
   ```bash
   python3 json-data.py
   ```

---

## File Mover

- **Files**: `mover.py`, `README.md` (within `mover/`)
- **Description**: Organizes and moves files based on their base names into corresponding subdirectories.
- **Usage**: Run the script in the target directory.
- **Run**: 
   ```bash
   python3 mover.py
   ```

---

## Renamer Script

- **Files**: `renamer.py`, `renamer.md`
- **Description**: Recursively renames files and directories to kebab case while excluding CamelCase and PascalCase.
- **Usage**: Run the script with the target directory.
- **Run**: 
   ```bash
   python3 renamer.py
   ```

---

## Speaker Formatter

- **Files**: `speaker.py`, `speaker.md`
- **Description**: Formats a transcription file by organizing speaker turns and timestamps for readability.
- **Usage**: Provide the transcription file as an argument.
- **Run**: 
   ```bash
   python3 speaker.py <input_markdown_file>
   ```

---

## Speedtest Logger

- **Files**: `speedtest.py`, `speedtest.md`
- **Description**: Logs network speed tests using `speedtest-cli` and stores the results in JSON format.
- **Usage**: Install `speedtest-cli` and run the script.
- **Run**: 
   ```bash
   python3 speedtest.py
   ```

---

## Status Bar Data

- **Files**: `status_bar.py`, `status_bar.md`
- **Description**: Collects and analyzes data to update a status bar.
- **Usage**: Run the script to gather and process data for the status bar.
- **Run**: 
   ```bash
   python3 status_bar.py
   ```

---

## Symlinks

- **Files**: `symlinks.sh`, `symlinks.md`
- **Description**: Creates symbolic links for specified files or directories to a target location.
- **Usage**: Define the files to be linked in the script and run it.
- **Run**: 
   ```bash
   ./symlinks.sh
   ```

---

## Today I Learned (TIL)

- **Files**: `til.sh`, `til.md`
- **Description**: A shell script to create and manage a daily learning log.
- **Usage**: Customize the script to suit your daily logging needs and run it.
- **Run**: 
   ```bash
   ./til.sh
   ```

---

## Virtual Environment

- **Files**: `venv/`
- **Description**: Contains a Python virtual environment for managing dependencies across scripts.
- **Usage**: Activate the virtual environment before running any Python scripts.
- **Run**: 
   ```bash
   source venv/bin/activate
   ```

---

## Watch Cronjobs

- **Files**: `watch_cronjobs.sh`, `watch_cronjobs.md`
- **Description**: A script to monitor the status of cron jobs by running a command or checking their output.
- **Usage**: Define the cron jobs or commands to monitor and run the script.
- **Run**: 
   ```bash
   ./watch_cronjobs.sh
   ```

---

## Word Check

- **Files**: `word_check.pl`, `word_check.py`, `word_check.md`
- **Description**: Filters and cleans a list of words based on criteria, removing duplicates

 and applying specific character rules.
- **Usage**:
   - **Perl**: 
     ```bash
     perl word_check.pl
     ```
   - **Python**:
     ```bash
     python3 word_check.py
     ```
- **Output**: Saves the cleaned word list to a file.

---

## License

All scripts in this directory are released under [The Unlicense](https://unlicense.org/), placing them in the public domain and free for unrestricted use.
