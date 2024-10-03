# Speedtest Logging Script

## Overview

This Python script runs a network speed test using `speedtest-cli` and logs the results (ping, download, and upload
speeds) in a JSON file. Each time the script is executed, it appends a new entry with the current timestamp and speed
test results to the file.

## Features

- **Automated Speed Test**: Runs a speed test using `speedtest-cli` and extracts `ping`, `download`, and `upload`
  results.
- **JSON Logging**: Appends the speed test results to a JSON file in a specified directory, maintaining a history of
  network performance over time.
- **Timestamped Entries**: Adds a timestamp to each entry for easy tracking of results over time.

## Usage

1. **Install `speedtest-cli`**: Make sure `speedtest-cli` is installed on your system. You can install it using `pip`:
   ```bash
   pip install speedtest-cli
   ```

2. **Run the Script**: Execute the script manually or schedule it to run periodically (e.g., using `cron` on Linux):
   ```bash
   python3 speedtest_logger.py
   ```

3. **Log Directory and File**:
    - The results are stored in `/home/jwl/Documents/speedtests/speedtest_results.json`. If the directory does not
      exist, it will be created automatically.
    - You can change the directory and file name by modifying the `directory` and `filename` variables at the top of the
      script.

## JSON File Structure

The JSON file contains an array of entries, each with:

- **timestamp**: The date and time of the speed test.
- **ping**: The latency in milliseconds (e.g., `"15 ms"`).
- **download**: The download speed in megabits per second (e.g., `"50 Mbit/s"`).
- **upload**: The upload speed in megabits per second (e.g., `"10 Mbit/s"`).

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

- **Change Output Directory or File**: Modify the `directory` and `filename` variables at the start of the script:
   ```python
   directory = "/your/desired/directory"
   filename = "your_results.json"
   ```
- **Error Handling**: The script checks for a JSON decoding error when loading the existing data. If the file is not a
  valid JSON, it starts with an empty array.

## Automation

To run the script automatically at regular intervals, you can use a scheduling tool like `cron`. For example, to run the
script every hour, add the following line to your crontab:

```bash
0 * * * * /usr/bin/python3 /path/to/speedtest_logger.py
```

## Requirements

- **Python 3**: The script is written in Python 3.
- **`speedtest-cli`**: Required to run the speed tests. Install via `pip install speedtest-cli`.
- **File System Access**: Ensure the script has write permissions to the specified directory.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.