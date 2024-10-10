# Update Crontab Script

## Overview

This Bash script updates the user's crontab using a predefined crontab file. It allows for easy management of cron jobs
by maintaining them in a version-controlled file, which can be applied to the system's crontab with a single command.

## Features

- **Automated Crontab Update**: Automatically sets the user's crontab to match the contents of a specified crontab file.
- **File Existence Check**: Verifies that the crontab file exists before attempting to update the crontab, providing an
  error message if the file is not found.
- **Centralized Crontab Management**: Makes it easy to version control and update cron jobs by storing them in a single
  file.

## Usage

1. **Set the Crontab File Path**: Make sure the `CRONTAB_FILE` variable in the script points to the correct path of your
   crontab file:
   ```bash
   CRONTAB_FILE="/path/to/your/cronjobs"
   ```

2. **Add or Modify Cron Jobs**: Edit the crontab file (e.g., `/home/jwl/dotfiles/.scripts/cronjobs`) to add or modify
   cron jobs as needed. The file should follow standard crontab syntax.

   **Example crontab file**:
   ```
   0 5 * * * /home/jwl/scripts/backup.sh
   */15 * * * * /home/jwl/scripts/check_updates.sh
   ```

3. **Run the Script**: Make sure the script is executable and then run it to update your crontab:
   ```bash
   chmod +x update_crontab.sh
   ./update_crontab.sh
   ```

   If the crontab file is found, the script will apply the cron jobs defined in that file to the user's crontab.

4. **Output**:
    - If the crontab file is found, the crontab will be updated without any further message.
    - If the crontab file is missing, the script will print an error message:
      ```
      Crontab file does not exist: /home/jwl/dotfiles/.scripts/cronjobs
      ```

## Requirements

- **Bash Shell**: The script is intended to run in a Bash shell.
- **Crontab Access**: Ensure you have permissions to update the user's crontab.
- **Valid Crontab File**: The crontab file must be in a valid format for cron to accept it.

## Customization

- **Change the Crontab File Path**: If your crontab file is located elsewhere, modify the `CRONTAB_FILE` variable:
  ```bash
  CRONTAB_FILE="/path/to/your/custom_cron_file"
  ```

## Notes

- **Error Handling**: The script checks for the existence of the crontab file before attempting to update the crontab.
  If the file does not exist, it will exit with an error.
- **Version Control**: By keeping your cron jobs in a file (e.g., within a Git repository), you can easily track changes
  and revert to previous versions if necessary.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.