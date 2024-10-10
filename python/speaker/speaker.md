# Transcription Formatting Script

## Overview

This Python script formats a transcription from a Markdown file, focusing on enhancing readability by organizing speaker
turns, timestamps, and text content. It processes the transcription part of the file, adds structure for speaker
changes, and outputs a formatted version of the transcription.

## Features

- **Speaker Turn Marking**: Splits the transcription based on `[SPEAKER_TURN]` markers to distinguish different
  speakers.
- **Timestamp Extraction**: Finds and organizes timestamps within the transcription.
- **Content Formatting**: Formats text for each timestamp by ensuring proper spacing and structure, removing unnecessary
  backslashes, and adding line breaks for clarity.

## Usage

1. **Run the Script**: Execute the script by passing the path to a Markdown file containing the transcription:
   ```bash
   ./speaker.py <input_markdown_file>
   ```
   Replace `<input_markdown_file>` with the path to your Markdown file.

   The script will create a new formatted Markdown file with `_formatted` appended to its name, e.g., `transcription.md`
   will become `transcription_formatted.md`.

2. **Output**: The formatted file will have:
    - Each `[SPEAKER_TURN]` on its own line.
    - Timestamps and corresponding text content organized for readability.
    - Removed unnecessary backslashes and added appropriate line breaks.

## Requirements

- **Python 3**: The script is written in Python 3.
- **Input File Format**: The script expects a Markdown file with a transcription section containing `[SPEAKER_TURN]`
  markers, timestamps (in brackets), and text content.

## Customization

- **Regex Modifications**: If your transcription format varies or requires different processing rules, adjust the
  `format_transcription()` function, particularly the regex patterns used to split the text and extract timestamps.

## Debugging Information

The script includes debugging print statements to help you understand how parts of the transcription are processed:

- **Processing Part**: Displays each segment being processed.
- **Matched Timestamps and Content**: Shows how timestamps and content are extracted and formatted.
- **Input and Output Debugging**: Prints both the original and formatted text for comparison.

To disable debugging output, simply remove or comment out the `print()` statements in the `format_transcription()` and
`main()` functions.

## Example

**Input File**: `transcription.md`

```
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

```
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