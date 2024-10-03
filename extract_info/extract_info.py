#!/usr/bin/env python3
import os
import re
import csv


def extract_info(email_content):
    # Extract estimated bill
    bill_match = re.search(
        r"Your next bill is estimated to be \$([0-9.]+)", email_content
    )
    estimated_bill = bill_match.group(1) if bill_match else None

    # Extract total weekly kWh usage
    kwh_match = re.search(r"Last week you used ([0-9]+) kWh", email_content)
    total_kwh = kwh_match.group(1) if kwh_match else None

    # Extract daily kWh usage
    daily_kwh_pattern = re.compile(
        r"Last week you used \d+ kWh\n((?:\d+\s*\n)+)Sun\tMon\tTue\tWed\tThu\tFri\tSat"
    )
    daily_kwh_match = daily_kwh_pattern.search(email_content)

    if daily_kwh_match:
        daily_kwh_values = daily_kwh_match.group(1).strip().split()
        daily_kwh = {
            "Sun": daily_kwh_values[0],
            "Mon": daily_kwh_values[1],
            "Tue": daily_kwh_values[2],
            "Wed": daily_kwh_values[3],
            "Thu": daily_kwh_values[4],
            "Fri": daily_kwh_values[5],
            "Sat": daily_kwh_values[6],
        }
    else:
        daily_kwh = {
            "Sun": "",
            "Mon": "",
            "Tue": "",
            "Wed": "",
            "Thu": "",
            "Fri": "",
            "Sat": "",
        }

    # Extract dates for the week
    dates_match = re.search(
        r"Subject: Your OG&E Energy Usage: (\d{2}/\d{2}/\d{4}) – (\d{2}/\d{2}/\d{4})",
        email_content,
    )
    week_start = dates_match.group(1) if dates_match else None
    week_end = dates_match.group(2) if dates_match else None

    return {
        "estimated_bill": estimated_bill,
        "total_kwh": total_kwh,
        "daily_kwh": daily_kwh,
        "week_start": week_start,
        "week_end": week_end,
    }


def process_directory(directory_path, output_file, debug_file):
    data = []
    debug_data = []

    # Iterate over all files in the directory
    for filename in os.listdir(directory_path):
        if filename.endswith(".txt"):  # Assuming the files are .txt, modify if needed
            file_path = os.path.join(directory_path, filename)

            with open(file_path, "r") as file:
                content = file.read()
                info = extract_info(content)
                data.append(info)
                debug_data.append({"filename": filename, "info": info})

    # Write data to a CSV file
    with open(output_file, "w", newline="") as csvfile:
        fieldnames = [
            "week_start",
            "week_end",
            "estimated_bill",
            "total_kwh",
            "Sun",
            "Mon",
            "Tue",
            "Wed",
            "Thu",
            "Fri",
            "Sat",
        ]
        writer = csv.DictWriter(csvfile, fieldnames=fieldnames)

        writer.writeheader()
        for row in data:
            row_data = {
                "week_start": row["week_start"],
                "week_end": row["week_end"],
                "estimated_bill": row["estimated_bill"],
                "total_kwh": row["total_kwh"],
                "Sun": row["daily_kwh"].get("Sun", ""),
                "Mon": row["daily_kwh"].get("Mon", ""),
                "Tue": row["daily_kwh"].get("Tue", ""),
                "Wed": row["daily_kwh"].get("Wed", ""),
                "Thu": row["daily_kwh"].get("Thu", ""),
                "Fri": row["daily_kwh"].get("Fri", ""),
                "Sat": row["daily_kwh"].get("Sat", ""),
            }
            writer.writerow(row_data)

    # Write debug information to a file
    with open(debug_file, "w") as debugfile:
        for entry in debug_data:
            debugfile.write(f"Filename: {entry['filename']}\n")
            for key, value in entry["info"].items():
                debugfile.write(f"{key}: {value}\n")
            debugfile.write("\n")


# Usage example
directory_path = os.getcwd()  # Use the current working directory
output_file = "collated_data.csv"
debug_file = "debug_log.txt"
process_directory(directory_path, output_file, debug_file)
