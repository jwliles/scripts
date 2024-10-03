#!/usr/bin/env python3

import json
import logging
import matplotlib.pyplot as plt

# Setup logging
logging.basicConfig(
    level=logging.INFO, format="%(asctime)s - %(levelname)s - %(message)s"
)


# Load the JSON data
def load_data(filename):
    try:
        with open(filename, "r") as file:
            data = json.load(file)
        logging.info(f"Data loaded from {filename}")
        return data
    except FileNotFoundError:
        logging.error(f"File not found: {filename}")
        return []
    except json.JSONDecodeError as e:
        logging.error(f"Error decoding JSON: {e}")
        return []


# Analyze the data
def analyze_data(data):
    download_speeds = []
    upload_speeds = []
    pings = []

    for entry in data:
        if (
            entry["download"] != "N/A Mbit/s"
            and entry["upload"] != "N/A Mbit/s"
            and entry["ping"] != "N/A ms"
        ):
            try:
                # Extract and convert speed from Mbit/s to float and ping from ms to float
                download_speeds.append(float(entry["download"].split()[0]))
                upload_speeds.append(float(entry["upload"].split()[0]))
                pings.append(float(entry["ping"].split()[0]))
            except ValueError as e:
                logging.error(f"Error parsing speeds or ping: {e}")

    average_download = (
        sum(download_speeds) / len(download_speeds) if download_speeds else 0
    )
    average_upload = sum(upload_speeds) / len(upload_speeds) if upload_speeds else 0
    min_ping = min(pings) if pings else 0
    max_ping = max(pings) if pings else 0

    logging.info(
        f"Analysis complete: Average Download {average_download:.2f} Mbps, Average Upload {average_upload:.2f} Mbps, Min Ping {min_ping} ms, Max Ping {max_ping} ms"
    )
    return average_download, average_upload, min_ping, max_ping


# Plot the data
def plot_data(data):
    timestamps = []
    download_speeds = []

    for entry in data:
        if entry["download"] != "N/A Mbit/s":
            try:
                timestamps.append(entry["timestamp"])
                download_speeds.append(float(entry["download"].split()[0]))
            except ValueError as e:
                logging.error(f"Error parsing download speed or timestamp: {e}")

    plt.figure()
    plt.plot(
        timestamps,
        download_speeds,
        marker="o",
        linestyle="-",
        color="b",
        label="Download Speeds",
    )
    plt.xlabel("Timestamp")
    plt.ylabel("Speed (Mbps)")
    plt.title("Download Speed Over Time")
    plt.xticks(rotation=45)  # Rotate timestamps for better readability
    plt.legend()
    plt.tight_layout()  # Adjust layout to make room for rotated x-axis labels
    plt.show()
    logging.info("Plot displayed successfully")


# Main function to run the script
def main():
    data = load_data("speedtest_results.json")
    if data:
        averages = analyze_data(data)
        print(f"Average Download Speed: {averages[0]:.2f} Mbps")
        print(f"Average Upload Speed: {averages[1]:.2f} Mbps")
        print(f"Minimum Ping: {averages[2]} ms")
        print(f"Maximum Ping: {averages[3]} ms")
        plot_data(data)
    else:
        logging.error("No data to process")


if __name__ == "__main__":
    main()
