#!/usr/bin/env python3

import os
import re
import sys

# List of Windows reserved names
WINDOWS_RESERVED_NAMES = [
    "CON",
    "PRN",
    "AUX",
    "NUL",
    "COM1",
    "COM2",
    "COM3",
    "COM4",
    "COM5",
    "COM6",
    "COM7",
    "COM8",
    "COM9",
    "LPT1",
    "LPT2",
    "LPT3",
    "LPT4",
    "LPT5",
    "LPT6",
    "LPT7",
    "LPT8",
    "LPT9",
]


def is_camel_or_pascal_case(name):
    # Check if the name is in CamelCase or PascalCase
    return bool(re.search(r"(?<!^)(?=[A-Z])", name))


def ensure_lowercase_extension(path):
    dirname, basename = os.path.split(path)
    name, ext = os.path.splitext(basename)
    return os.path.join(dirname, name + ext.lower())


def rename_item(path, is_dir=False):
    dirname, basename = os.path.split(path)
    name, ext = os.path.splitext(basename)

    # Special handling for README.md to be all uppercase
    if basename.lower() == "readme.md":
        new_name = "README" + ext.upper()
    elif is_camel_or_pascal_case(name):
        # Skip renaming CamelCase or PascalCase files and directories
        return path
    else:
        # Convert name to lowercase
        new_name = name.lower()

        # Replace "&" with "and"
        new_name = new_name.replace("&", "and")

        # Remove quotes, apostrophes, and punctuation
        new_name = re.sub(r"[\"\'.,!?\(\)]", "", new_name)

        # Replace colons, en dashes, em dashes, and hyphens with underscores
        new_name = re.sub(r"[:–—-]", "_", new_name)

        # Replace Windows-incompatible characters with underscores
        new_name = re.sub(r'[<>:"/\\|?*]', "_", new_name)

        # Replace multiple underscores with a single underscore
        new_name = re.sub(r"_+", "_", new_name)

        # Remove leading or trailing underscores
        new_name = new_name.strip("_")

        # Handle reserved names by appending an underscore
        if new_name.upper() in WINDOWS_RESERVED_NAMES:
            new_name += "_"

        if not is_dir:
            # Combine the new name with the original extension in lowercase
            new_name += ext.lower()

    new_path = os.path.join(dirname, new_name)

    if new_path != path:
        os.rename(path, new_path)

    return new_path


def rename_files_and_directories(directory):
    for root, dirs, files in os.walk(directory, topdown=False):
        # Exclude hidden directories
        dirs[:] = [d for d in dirs if not d.startswith(".")]

        for filename in files:
            if filename.startswith("."):
                continue
            old_path = os.path.join(root, filename)
            lowercased_path = ensure_lowercase_extension(old_path)
            rename_item(lowercased_path)

        for dirname in dirs:
            if dirname.startswith("."):
                continue
            old_path = os.path.join(root, dirname)
            rename_item(old_path, is_dir=True)


if __name__ == "__main__":
    # Get the directory from the command line argument, or use the current working directory
    directory = sys.argv[1] if len(sys.argv) > 1 else os.getcwd()
    rename_files_and_directories(directory)
