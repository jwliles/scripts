# Git Repository Cloning Script

## Overview

This Bash script clones all Git repositories in the current directory to a specified target directory. It scans for all
`.git` directories, extracts their parent folder names as project names, and clones each repository to the target
location.

## Features

- **Bulk Repository Cloning**: Automatically identifies and clones all Git repositories in the current directory.
- **Custom Target Directory**: Allows the user to specify a target directory for cloned repositories.
- **Automatic Directory Creation**: Creates the target directory if it doesn't already exist.

## Usage

1. **Run the Script**: Use the following command to execute the script, replacing `<target-directory>` with your desired
   directory:
   ```bash
   ./clone_repos.sh <target-directory>
   ```
    - `<target-directory>`: The path where you want all repositories to be cloned.

2. The script will loop over all `.git` directories in the current directory, clone each repository, and store them in
   `<target-directory>`.

## Example

```bash
./clone_repos.sh ~/my_repos
```

This will clone all Git repositories from the current directory to `~/my_repos`.

## Requirements

- **Bash**: The script should be run in a Bash shell.
- **Git**: Ensure that `git` is installed and accessible in your `PATH`.

## Customization

- **Target Directory**: Modify the script to change the default target directory or add additional logic for more
  complex directory structures.
- **Repository Detection**: The script assumes all subdirectories with `.git` are Git repositories. If you have a
  different structure, you may need to adjust the repository detection logic.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.