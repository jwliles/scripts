# Dotforge (df)

**Dotforge** is a command-line tool designed to efficiently manage symlinks, track approved file types, and configure ignored paths with ease. With both short and long command options, Dotforge streamlines your workflow by automating repetitive tasks related to directory and file management.

---

## Installation

### Prerequisites

- **Go** (Golang) installed on your system.

### Build from Source

Clone the repository:

```bash
git clone https://github.com/your-username/dotforge.git
cd dotforge
```

Build the binary:

```bash
go build -o df
```

Move `df` to a directory in your `PATH`:

```bash
mv df /usr/local/bin/ # or any directory in your PATH
```

### Verify Installation

```bash
df --help
```

---

## Usage

Dotforge provides both long and short options for all commands. Below is a breakdown of all available options:

### General Commands

#### Set the Default Path

- **Command**: `--set-default-path`, `-sdp <path>`
- **Description**: Sets the default target directory where symlinks will be created.
- **Example**:

  ```bash
  df --set-default-path /home/user/Applications
  df -sdp /home/user/Applications
  ```

#### Add New File Types

- **Command**: `--new-filetypes`, `-nft <extensions>`
- **Description**: Adds file types to the approved list for symlinking (e.g., `.go`, `.py`).
- **Example**:

  ```bash
  df --new-filetypes .go .sh
  df -nft .go .sh
  ```

#### Remove File Types

- **Command**: `--remove-filetypes`, `-rft <extensions>`
- **Description**: Removes file types from the approved list.
- **Example**:

  ```bash
  df --remove-filetypes .txt
  df -rft .txt
  ```

#### List All File Types

- **Command**: `--list-filetypes`, `-lft`
- **Description**: Displays all approved file types.
- **Example**:

  ```bash
  df --list-filetypes
  df -lft
  ```

#### Add Ignored Paths

- **Command**: `--ignore-paths`, `-ip <paths>`
- **Description**: Adds paths (files or directories) to the ignored list.
- **Example**:

  ```bash
  df --ignore-paths /home/user/temp
  df -ip /home/user/temp
  ```

#### Remove Ignored Paths

- **Command**: `--remove-ignored`, `-ri <paths>`
- **Description**: Removes paths from the ignored list.
- **Example**:

  ```bash
  df --remove-ignored /home/user/temp
  df -ri /home/user/temp
  ```

#### List All Ignored Paths

- **Command**: `--list-ignored`, `-li`
- **Description**: Displays all paths currently ignored by Dotforge.
- **Example**:

  ```bash
  df --list-ignored
  df -li
  ```

#### Scan the Current Directory

- **Command**: `--scan`, `-s`
- **Description**: Scans the current directory for files that should be symlinked or ignored.
- **Example**:

  ```bash
  df --scan
  df -s
  ```

#### Set Max Lines for Display

- **Command**: `--max-lines`, `-ml <number>`
- **Description**: Sets the maximum number of lines to display before output is opened in an editor (default: 20).
- **Example**:

  ```bash
  df --scan --max-lines 50
  df -s -ml 50
  ```

#### List the Default Path

- **Command**: `--list-default-path`, `-ldp`
- **Description**: Displays the current default target path.
- **Example**:

  ```bash
  df --list-default-path
  df -ldp
  ```

#### Show Help

- **Command**: `--help`, `-h`
- **Description**: Displays help information for Dotforge.
- **Example**:

  ```bash
  df --help
  df -h
  ```

## Example Usage

- **Set a Default Path**:

  ```bash
  df -sdp /home/user/Applications
  ```

- **Add File Types**:

  ```bash
  df -nft .go .py
  ```

- **Ignore Paths**:

  ```bash
  df -ip /home/user/temp /home/user/logs
  ```

- **Scan Directory**:

  ```bash
  df -s
  ```

---

## Troubleshooting

### Common Issues

#### Command Not Found

- **Issue**: `command not found: df`
- **Solution**: Make sure `df` is in your `PATH` and is executable. If not, add it:

  ```bash
  export PATH=$PATH:/path/to/df
  chmod +x /path/to/df
  ```

#### Permission Denied

- **Issue**: `permission denied` when running `df`.
- **Solution**: Check the permissions of the target directory and ensure you have the necessary rights.

---

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes.

---

## License

This project is licensed under [The Unlicense](https://unlicense.org/).

---

## Acknowledgments

Special thanks to all contributors and the open-source community for their continuous support and feedback.
