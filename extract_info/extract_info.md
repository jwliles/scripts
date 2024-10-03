# Energy Usage Data Extraction Script

## Overview

This Python script processes a collection of text files containing energy usage data, extracting relevant details such
as estimated bills, total weekly kWh usage, and daily kWh usage. The script then consolidates this data into a CSV file
for easy analysis and also logs detailed debug information for verification.

## Features

- **Data Extraction**: Parses text files to extract:
    - Estimated bill amounts.
    - Total weekly energy usage in kWh.
    - Daily energy usage for each day of the week.
    - Start and end dates of the usage week.
- **CSV Output**: Aggregates the extracted data into a structured CSV file.
- **Debug Log**: Logs detailed extraction information to a debug file for troubleshooting.

## Usage

1. **Directory Setup**: Place all text files (`.txt`) containing energy usage data in a directory that the script will
   process. By default, the script processes files in the current working directory.

2. **Run the Script**:
   ```bash
   python3 energy_usage_extractor.py
   ```

   The script will:
    - Extract data from all `.txt` files in the directory.
    - Generate a CSV file (`collated_data.csv`) with the extracted data.
    - Create a debug log file (`debug_log.txt`) containing detailed information about the extracted content.

3. **CSV Output**: The resulting CSV file (`collated_data.csv`) will contain the following columns:
    - `week_start`: Start date of the usage week.
    - `week_end`: End date of the usage week.
    - `estimated_bill`: Estimated bill amount in dollars.
    - `total_kwh`: Total energy usage in kWh for the week.
    - `Sun`, `Mon`, `Tue`, `Wed`, `Thu`, `Fri`, `Sat`: Daily kWh usage for each day of the week.

## Requirements

- **Python 3**: The script is written in Python 3 and requires standard libraries (`os`, `re`, `csv`).
- **Text Files**: Ensure the input files are in `.txt` format and contain energy usage details in the expected
  structure.

## Customization

- **Directory Path**: Modify `directory_path` to specify a different directory for the input files:
  ```python
  directory_path = "/path/to/your/directory"
  ```

- **Output File Names**: Change `output_file` and `debug_file` to customize the names of the generated CSV and debug log
  files:
  ```python
  output_file = "your_output_file.csv"
  debug_file = "your_debug_file.txt"
  ```

## Debug Information

The script logs extracted data from each file in `debug_log.txt`, which includes:

- Filename processed.
- Extracted details such as estimated bill, total kWh usage, and daily kWh usage.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.