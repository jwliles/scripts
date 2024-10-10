# Directory Organizer Script

## Overview

This script scans a directory, organizes files by their base names, and moves them into corresponding folders. If a file shares the same base name (ignoring extensions) with other files, all such files are moved into a newly created folder named after their base name. The script skips all hidden files and folders by default (those starting with a `.`).

For example, if the directory contains `example.py`, `example.md`, and `example.sh`, the script will create a folder named `example` and move all three files into this folder.

## Usage

1. **Ensure the script is executable**:
   You may need to change the permissions to make the script executable:
   ```bash
   chmod +x your_script_name.py


## License
These scripts are open-source and available for free use and modification under [The Unlicense](https://unlicense.org/), making them public domain and free to use without any restrictions.
