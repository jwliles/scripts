Certainly! Below are a comprehensive **README** and a **man page** for the current version of the `pathfinder` tool. These documents will help users understand how to install, use, and benefit from `pathfinder`.

---

# **README**

# pathfinder

**Version:** 1.0  
**Author:** Your Name  
**License:** MIT License

## **Table of Contents**

- [**README**](#readme)
- [pathfinder](#pathfinder)
  - [**Table of Contents**](#table-of-contents)
  - [**Introduction**](#introduction)
  - [**Features**](#features)
  - [**Installation**](#installation)
    - [**Prerequisites**](#prerequisites)
    - [**Building from Source**](#building-from-source)
  - [**Usage**](#usage)
    - [**Command Overview**](#command-overview)
    - [**Commands**](#commands)
      - [**add**](#add)
      - [**remove**](#remove)
      - [**list**](#list)
      - [**history**](#history)
      - [**restore**](#restore)
    - [**Examples**](#examples)
  - [**Configuration**](#configuration)
  - [**Backup Management**](#backup-management)
  - [**Contributing**](#contributing)
  - [**License**](#license)

## **Introduction**

**pathfinder** is a powerful command-line tool written in Rust for managing your system's `PATH` environment variable. It simplifies the process of adding and removing directories from your `PATH`, ensures backups are created automatically, and provides tools to restore previous configurations.

Managing the `PATH` variable is crucial for system performance and command execution. `pathfinder` provides a safe and efficient way to handle `PATH` modifications, with features designed to prevent errors and maintain system stability.

## **Features**

- **Effortless Management**: Easily add or remove directories from your `PATH`.
- **Automatic Backups**: Creates time-stamped backups of your `PATH` before any changes.
- **Restoration**: Restore your `PATH` from any previous backup.
- **Listing**: View all current entries in your `PATH`.
- **Cross-Platform**: Compatible with Unix/Linux and macOS systems.
- **Safe Modifications**: Validates directories before adding them to prevent errors.
- **Persistent Changes**: Updates your shell configuration to make changes permanent.

## **Installation**

### **Prerequisites**

- **Rust Toolchain**: Ensure you have Rust installed. You can install Rust using [rustup](https://www.rust-lang.org/tools/install):

  ```bash
  curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
  ```

### **Building from Source**

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/pathfinder.git
   cd pathfinder
   ```

2. **Build the Project**

   ```bash
   cargo build --release
   ```

   This command compiles the project in release mode, producing an optimized binary.

3. **Install the Binary**

   Optionally, you can install the binary system-wide:

   ```bash
   sudo cp target/release/pathfinder /usr/local/bin/
   ```

   Or add it to your `PATH`:

   ```bash
   export PATH="$PATH:$(pwd)/target/release"
   ```

## **Usage**

### **Command Overview**

```bash
pathfinder [COMMAND] [OPTIONS]
```

### **Commands**

#### **add**

Add a directory to your `PATH`.

**Usage:**

```bash
pathfinder add <directory>
```

**Options:**

- `<directory>`: The directory path to add to your `PATH`.

#### **remove**

Remove a directory from your `PATH`.

**Usage:**

```bash
pathfinder remove <directory>
```

**Options:**

- `<directory>`: The directory path to remove from your `PATH`.

#### **list**

List all current entries in your `PATH`.

**Usage:**

```bash
pathfinder list
```

#### **history**

Show the backup history of your `PATH`.

**Usage:**

```bash
pathfinder history
```

#### **restore**

Restore your `PATH` from a previous backup.

**Usage:**

```bash
pathfinder restore [--timestamp <timestamp>]
```

**Options:**

- `--timestamp <timestamp>`: (Optional) The timestamp of the backup to restore. If not provided, the most recent backup is used.

### **Examples**

- **Add a Directory to PATH**

  ```bash
  pathfinder add ~/my/custom/bin
  ```

- **Remove a Directory from PATH**

  ```bash
  pathfinder remove ~/my/old/bin
  ```

- **List PATH Entries**

  ```bash
  pathfinder list
  ```

  **Sample Output:**

  ```
  Current PATH entries:
  - /usr/local/bin
  - /usr/bin
  - /bin
  - /usr/local/sbin
  - /usr/sbin
  - /sbin
  - ~/my/custom/bin
  ```

- **Show Backup History**

  ```bash
  pathfinder history
  ```

  **Sample Output:**

  ```
  Available backups:
  - backup_20231007_120000.json
  - backup_20231008_090000.json
  ```

- **Restore PATH from a Specific Backup**

  ```bash
  pathfinder restore --timestamp 20231007_120000
  ```

- **Restore PATH from the Most Recent Backup**

  ```bash
  pathfinder restore
  ```

## **Configuration**

`pathfinder` modifies your shell configuration file to make changes to `PATH` persistent across sessions.

- **Supported Shells**: Bash (`.bashrc`), Zsh (`.zshrc`), or a generic `.profile` if the shell is not recognized.
- **Backup Directory**: Backups are stored in `~/.pathfinder_backups`.

**Note**: Always review changes made to your shell configuration files. `pathfinder` adds an export command to update your `PATH`.

## **Backup Management**

- **Automatic Backups**: Before any modification, `pathfinder` creates a backup of your current `PATH` with a timestamp.
- **Backup Files**: Stored as JSON files in `~/.pathfinder_backups`.
- **Restoration**: Use the `restore` command to revert to a previous `PATH` configuration.

**Backup File Format Example (`backup_20231008_090000.json`):**

```json
{
  "timestamp": "20231008_090000",
  "path": "/usr/local/bin:/usr/bin:/bin:~/my/custom/bin"
}
```

## **Contributing**

Contributions are welcome! Please follow these steps:

1. **Fork the Repository**
2. **Create a Feature Branch**

   ```bash
   git checkout -b feature/your-feature
   ```

3. **Commit Your Changes**

   ```bash
   git commit -am 'Add your feature'
   ```

4. **Push to the Branch**

   ```bash
   git push origin feature/your-feature
   ```

5. **Create a Pull Request**

## **License**

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
