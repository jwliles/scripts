# Word Check Script

## Overview

These scripts, one written in Perl and the other in Python, process a text file containing words, filter them based on
specific criteria, and output a cleaned list of words. The main purpose of the scripts is to remove duplicates, words
containing certain patterns, and words that don't meet specific length or character requirements.

## Features

- **Converts Words to Lowercase**: All words in the input file are converted to lowercase for uniformity.
- **Removes Punctuation**: Any punctuation within the words is stripped before processing.
- **Filters Words Based on Criteria**:
    - Must contain only English letters (A-Z, case-insensitive).
    - Must be at least 4 characters long.
    - Must not have three or more consecutive identical letters (e.g., "aaa").
    - Must not start or end with two or more identical letters (e.g., "aab", "baa").
- **Ensures Unique Words**: No duplicate words are included in the output.

## Input and Output

- **Input**: Both scripts read from an input file named `word_list.txt`.
- **Output**: The filtered words are saved to:
    - **Perl Script**: `clean_word_list.txt`
    - **Python Script**: `clean_list.txt`

## How to Use

### Perl Script

1. Place your word list in a file named `word_list.txt` in the same directory as the script.
2. Run the script:
   ```bash
   perl word_check.pl
   ```
3. The filtered and cleaned word list will be written to `clean_word_list.txt`.

### Python Script

1. Place your word list in a file named `word_list.txt` in the same directory as the script.
2. Run the script:
   ```bash
   python3 clean_word_list.py
   ```
3. The filtered and cleaned word list will be written to `clean_list.txt`.

## Requirements

### For the Perl Script

- **Perl**: Ensure Perl is installed on your system.
- **UTF-8 Encoding**: The input file should be UTF-8 encoded for proper processing.

### For the Python Script

- **Python 3**: The script is written for Python 3 and uses standard libraries (`string`, `re`).

## Customization

To change the input or output file names, modify the relevant lines in either script:

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

Set `input_file` and `output_file` to any file paths of your choice.

## Differences Between Perl and Python Versions

While both scripts achieve the same purpose, there are some differences in implementation:

- **Function Organization**: The Python script uses separate functions to modularize checks for word validity, whereas
  the Perl script handles all checks within a single loop.
- **Output File Name**: The output file names are different by default (`clean_word_list.txt` for Perl and
  `clean_list.txt` for Python).

## License

These scripts are open-source and available for free use and modification under [The Unlicense](https://unlicense.org/),
making them public domain and free to use without any restrictions.