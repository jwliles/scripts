#!/usr/bin/env python3

import os
import re
import argparse


def is_empty_file(content):
    # Regex pattern to match YAML frontmatter
    yaml_pattern = re.compile(r"^---\s*\n(.*?)\n---\s*\n", re.DOTALL)
    # Regex pattern to match H1 heading possibly with newlines in between
    h1_pattern = re.compile(r"^\s*# .*\n?$")

    # Check for YAML frontmatter
    yaml_match = yaml_pattern.match(content)
    if yaml_match:
        # Get the rest of the content after YAML frontmatter
        rest_of_content = content[yaml_match.end() :].strip()
        # Check if the rest of the content is empty or only an H1 heading
        if not rest_of_content or (
            len(rest_of_content.splitlines()) == 1 and h1_pattern.match(rest_of_content)
        ):
            return True
    return False


def find_and_delete_empty_files(directory, dry_run=False):
    for root, _, files in os.walk(directory):
        for file in files:
            if file.endswith(".md"):
                file_path = os.path.join(root, file)
                with open(file_path, "r", encoding="utf-8") as f:
                    content = f.read()
                    if is_empty_file(content):
                        if dry_run:
                            print(f"Would delete {file_path}")
                        else:
                            print(f"Deleting {file_path}")
                            os.remove(file_path)


if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        description="Delete Markdown files with only YAML frontmatter and an H1 heading."
    )
    parser.add_argument(
        "--dry-run",
        action="store_true",
        help="Perform a dry run without deleting any files",
    )
    args = parser.parse_args()

    current_directory = os.getcwd()  # Get the current working directory
    find_and_delete_empty_files(current_directory, dry_run=args.dry_run)
