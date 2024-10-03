# Cron Jobs Overview

## Overview

This README provides a summary of the cron jobs configured for automation tasks on your system. These jobs perform
regular file backups, speed tests, and data fetching, each running at a specified time of day.

## Cron Job Schedule

1. **Backup `.zsh_history` File**
    - **Time**: Every hour at the 5th minute.
    - **Command**:
      ```bash
      cp /home/jwl/.zsh_history /home/jwl/dotfiles/
      ```
    - **Description**: Copies the `.zsh_history` file to the `dotfiles` directory for backup and versioning.

2. **Run Speed Test Script**
    - **Time**: Every hour at the 6th minute.
    - **Command**:
      ```bash
      /home/jwl/.scripts/speedtest.py
      ```
    - **Description**: Executes a Python script to perform a speed test. The script is located in `.scripts` directory.

3. **Log Speed Test Results**
    - **Time**: Every hour at the 35th minute.
    - **Command**:
      ```bash
      /usr/bin/python3 /home/jwl/projects/speedtests/python/speedtest.py >> /home/jwl/projects/speedtests/logs/speedtest.log 2>&1
      ```
    - **Description**: Runs the Python speed test script and appends the output (including errors) to `speedtest.log`.

4. **Fetch Speed Test Data**
    - **Time**: Every hour at the start of the hour.
    - **Command**:
      ```bash
      /home/jwl/projects/speedtests/frontend/speedtest-results/fetch-data.sh >> /home/jwl/projects/speedtests/logs/fetch-data.log 2>&1
      ```
    - **Description**: Runs a shell script to fetch speed test data and appends the output (including errors) to
      `fetch-data.log`.

## Crontab Format Reference

```
<Minute> <Hour> <Day of Month> <Month> <Day of Week> <Command>
```

For all the cron jobs above:

- **Minute**: The minute of each hour when the job is run.
- **Hour**: The hour of the day (24-hour format) when the job is run.
- **Command**: The specific command to execute, including any redirection of output.

## Notes

- **Logging**: Two of the jobs redirect their output and errors to log files in the `logs` directory.
- **Script Paths**: Make sure all script paths (`speedtest.py`, `fetch-data.sh`) are correct and the scripts have
  execute permissions.

## License

This script is released under [The Unlicense](https://unlicense.org/), making it public domain and free to use without
any restrictions.