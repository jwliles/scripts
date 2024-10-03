#!/usr/bin/env python3

import re
import sys


def format_transcription(text):
    # Remove all unnecessary backslashes
    text = text.replace("\\", "")

    # Split the text by speaker turns
    parts = re.split(r"(\[SPEAKER_TURN\])", text)

    formatted_text = ""
    for part in parts:
        # Strip leading/trailing whitespace from each part
        part = part.strip()

        # Debugging: Print the current part being processed
        print(f"Processing part: {part}\n")

        # Add speaker turn markers
        if part == "[SPEAKER_TURN]":
            formatted_text += f"{part}\n"
            continue

        # Find timestamps and text within each part
        matches = re.findall(r"(\[.*?\])\s*(.*?)(?=(\[.*?\])|\Z)", part, re.DOTALL)
        for match in matches:
            timestamp = match[0]
            content = match[1].replace("\n", " ").strip()
            # Debugging: Print the matched timestamp and content
            print(f"Matched timestamp: {timestamp}\nMatched content: {content}\n")
            formatted_text += f"{timestamp}\n{content}\n\n"

    # Ensure the text ends with a single newline
    formatted_text = formatted_text.strip() + "\n"
    return formatted_text


def main():
    if len(sys.argv) != 2:
        print("Usage: ./speaker.py <input_markdown_file>")
        sys.exit(1)

    input_filename = sys.argv[1]
    output_filename = input_filename.replace(".md", "_formatted.md")

    # Read input markdown file
    with open(input_filename, "r") as file:
        text = file.read()

    # Extract transcription part after metadata
    transcription_start = text.find("\n\n") + 2
    transcription_text = text[transcription_start:]

    # Print the input text for debugging
    print("Input text:")
    print(text)
    print("\n" + "=" * 40 + "\n")

    # Format the transcription
    formatted_text = format_transcription(transcription_text)

    # Print the formatted text for debugging
    print("Formatted text:")
    print(formatted_text)

    # Write to output markdown file
    with open(output_filename, "w") as file:
        file.write(formatted_text)


if __name__ == "__main__":
    main()
