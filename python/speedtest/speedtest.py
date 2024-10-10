#!/usr/bin/env python3

import subprocess
import datetime
import json
import os

# Directory and file setup
directory = "/home/jwl/Documents/speedtests"
filename = "speedtest_results.json"
output_file = os.path.join(directory, filename)

# Ensure the directory exists
os.makedirs(directory, exist_ok=True)


# Function to run speedtest and save results in JSON format
def run_speedtest():
    # Run speedtest-cli
    result = subprocess.run(
        ["speedtest-cli", "--simple"], capture_output=True, text=True
    )
    results_str = result.stdout

    # Extract ping, download, and upload results
    ping = (
        next(
            (line.split(" ")[1] for line in results_str.split("\n") if "Ping" in line),
            "N/A",
        )
        + " ms"
    )
    download = (
        next(
            (
                line.split(" ")[1]
                for line in results_str.split("\n")
                if "Download" in line
            ),
            "N/A",
        )
        + " Mbit/s"
    )
    upload = (
        next(
            (
                line.split(" ")[1]
                for line in results_str.split("\n")
                if "Upload" in line
            ),
            "N/A",
        )
        + " Mbit/s"
    )

    # Current timestamp
    now = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")

    # Prepare the entry
    entry = {"timestamp": now, "ping": ping, "download": download, "upload": upload}

    # Check if file exists and load existing data if it does
    if os.path.exists(output_file):
        with open(output_file, "r") as f:
            try:
                data = json.load(f)
            except json.JSONDecodeError:
                data = []
    else:
        data = []

    # Append the new entry
    data.append(entry)

    # Save the updated data to the file
    with open(output_file, "w") as f:
        json.dump(data, f, indent=4)


run_speedtest()
