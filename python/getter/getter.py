#!/usr/bin/env python3

import os


def get_script_info(file_path):
    info = {}
    description_lines = []

    with open(file_path, "r") as file:
        lines = file.readlines()
        if lines:
            # Extract shebang if present
            first_line = lines[0].strip()
            info["shebang"] = first_line if first_line.startswith("#!") else "Unknown"

            # Collect single-line comments from the file
            for line in lines[1:]:  # Skip shebang line
                stripped_line = line.strip()
                # Handle only single-line comments that start with `#` or `//`
                if stripped_line.startswith("#") or stripped_line.startswith("//"):
                    # Remove leading comment characters and extra whitespace
                    clean_line = stripped_line.lstrip("#/ ").strip()
                    # Wrap each comment line in single backticks for inline code formatting
                    description_lines.append(f"{clean_line}  \n")

    # Join lines with spaces to keep comments in the same paragraph
    info["description"] = " ".join(description_lines)
    return info


def generate_readme(directory):
    readme_content = "# Script Index\n\n"
    print(f"Scanning directory: {directory}")

    # List only files in the root directory (no recursion)
    for file in os.listdir(directory):
        file_path = os.path.join(directory, file)
        if os.path.isfile(file_path) and file.endswith(
            (".sh", ".py", ".go", ".rs", ".pl", ".rb")
        ):
            script_info = get_script_info(file_path)

            readme_content += f"## {file}\n\n"
            readme_content += f"**Path**: `{file_path}`  \n"
            readme_content += (
                f"**Shebang**: `{script_info.get('shebang', 'Unknown')}`  \n"
            )
            # Add description lines as inline code
            readme_content += f"**Description**:\n{script_info.get('description', 'No description').strip()}  \n\n\n"
            readme_content += "---\n\n"

    # Write the collected content to a README.md file
    readme_path = os.path.join(directory, "README.md")
    with open(readme_path, "w") as readme_file:
        readme_file.write(readme_content)
    print(f"README.md created at {readme_path}")


if __name__ == "__main__":
    generate_readme(".")
