# Speed Test Data Analyzer

## Overview

This Python script loads speed test results from a JSON file, analyzes the data to compute average download/upload
speeds and ping times, and visualizes the download speed trend over time using a line plot.

## Features

- **Load Speed Test Data**: Reads speed test results from a JSON file.
- **Data Analysis**: Computes:
    - Average download speed (Mbps).
    - Average upload speed (Mbps).
    - Minimum and maximum ping times (ms).
- **Visualization**: Plots download speeds over time to provide a visual representation of speed trends.

## Usage

1. **Prepare Your JSON Data**: Ensure you have a JSON file named `speedtest_results.json` in the same directory as the
   script, or modify the `load_data` function to use a different file path. The JSON structure should be an array of
   entries containing `timestamp`, `download`, `upload`, and `ping` fields, for example:
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

2. **Install Dependencies**: This script requires `matplotlib` for plotting:
   ```bash
   pip install matplotlib
   ```

3. **Run the Script**:
   ```bash
   python3 speedtest_analyzer.py
   ```

   The script will:
    - Load the JSON data.
    - Analyze the data and print the average download/upload speeds and min/max ping.
    - Display a plot of download speeds over time.

## Requirements

- **Python 3**: The script is written in Python 3.
- **`matplotlib`**: Used for plotting download speed trends. Install using `pip`.
- **JSON Data File**: Ensure the JSON data is correctly formatted and contains the required fields.

## Customization

- **Change Data File**: If your data is not in `speedtest_results.json`, modify the `load_data` function to point to the
  correct file:
  ```python
  data = load_data("/path/to/your/data.json")
  ```

- **Adjust Data Parsing**: The script expects download and upload speeds in "Mbit/s" and ping times in "ms". Adjust the
  parsing in the `analyze_data` function if your format differs.

## Notes

- **Error Handling**: The script logs errors for file handling, JSON decoding issues, and value parsing problems, making
  it easier to debug data inconsistencies.
- **Data Visualization**: The script currently only plots download speeds. You can extend the `plot_data` function to
  plot other metrics such as upload speed or ping times.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.