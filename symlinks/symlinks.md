# Emacs Configuration Symlink Script

## Overview

This Bash script creates symbolic links (symlinks) from an Emacs configuration directory within a dotfiles repository to
the user's home `.emacs.d` directory. This approach makes it easy to manage and version control Emacs configuration
files.

## Features

- **Automatic Directory Creation**: Ensures that the target `.emacs.d` directory exists in the user's home directory.
- **Symlink Creation**: Creates symbolic links for specific files or directories from the source dotfiles directory to
  the target Emacs configuration directory.
- **Easily Manageable Configuration**: Allows for easy management of Emacs configuration files through symlinks, making
  updates and backups more efficient.

## Usage

1. **Define the Files to be Symlinked**:
    - Modify the `FILES` array in the script to include the names of the files and directories you want to symlink:
      ```bash
      FILES=("init.el" "lisp" "config.org")
      ```
      Each entry should match the relative path of a file or directory inside your `dotfiles/.emacs.d` directory.

2. **Run the Script**:
    - Make sure the script is executable:
      ```bash
      chmod +x symlink_emacs.sh
      ```
    - Execute the script to create the symlinks:
      ```bash
      ./symlink_emacs.sh
      ```

   The script will create symlinks for each file/directory in the `FILES` array from `~/dotfiles/.emacs.d` to
   `~/.emacs.d`.

## Example

**Directory Structure**:

- Source directory (`~/dotfiles/.emacs.d/`):
  ```
  init.el
  lisp/
  config.org
  ```

- `FILES` array:
  ```bash
  FILES=("init.el" "lisp" "config.org")
  ```

After running the script, your `~/.emacs.d/` directory will contain symlinks pointing to the corresponding files and
directories in `~/dotfiles/.emacs.d/`.

## Requirements

- **Bash Shell**: The script is designed to be run in a Bash shell.
- **Read and Write Permissions**: Ensure that you have the necessary permissions to create directories and symlinks in
  the target location.

## Customization

- **Source and Target Directories**: Modify `SOURCE_DIR` and `TARGET_DIR` if your dotfiles or Emacs configuration are
  located in different directories:
  ```bash
  SOURCE_DIR="/path/to/your/dotfiles/.emacs.d"
  TARGET_DIR="/path/to/your/.emacs.d"
  ```

- **Symlink Files and Directories**: Adjust the `FILES` array to include all files and directories you wish to symlink.

## Notes

- **Avoid Overwriting Existing Files**: If the target files already exist, the script may fail to create the symlinks.
  Make sure to backup or remove existing files if necessary.
- **Safe to Run Multiple Times**: If new files are added to `FILES`, rerun the script to create symlinks for the new
  additions.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.