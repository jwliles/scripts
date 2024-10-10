# Directory Structure Generator

## Overview

This Go script generates a customizable directory structure with a specified number of folders and files. It provides different levels of complexity, from low to extreme, and allows for custom settings as well. The script creates files with random content and various extensions, ensuring that no directory is left empty.

## Features

- **Configurable Complexity Levels**: Offers preset complexity levels:
  - **Low**: 100 folders, 500 files.
  - **Medium**: 1000 folders, 5000 files.
  - **High**: 2000 folders, 10,000 files.
  - **Extreme**: 5000 folders, 200,000 files.
- **Custom Configuration**: Allows users to define their own folder and file limits.
- **Random File Generation**: Creates files with randomly chosen extensions and random content.
- **Ensures Non-Empty Directories**: Each directory is guaranteed to contain at least one file.
- **Parallel Directory and File Creation**: Uses Go routines to efficiently create directories and files in parallel.

## Usage

1. **Compile and Run the Script**: Make sure you have Go installed. Then compile and run the script with:
   ```bash
   go run directory_generator.go <target_directory>
   ```
   Replace `<target_directory>` with the path where you want the structure to be generated.

2. **Choose Complexity Level**: Upon running the script, you will be prompted to select a complexity level:
   ```
   Choose a directory complexity level:
     1. Low (100 folders, 500 files)
     2. Medium (1000 folders, 5000 files)
     3. High (2000 folders, 10,000 files)
     4. Extreme (5000 folders, 200,000 files)
     5. Custom (Enter your own values)
   ```
   Select the appropriate level or choose the custom option to input your own folder and file limits.

## Requirements

- **Go Language**: The script is written in Go and requires a Go environment to compile and run.
- **File System Access**: The script will create files and directories, so make sure you have the appropriate permissions in the target directory.

## Customization

- **File Extensions**: The script uses a set of predefined file extensions, which are weighted to approximate a realistic dataset:
  ```go
  var fileExtensions = []string{"md", "nomedia", "png", "jpg", "json", ...}
  ```
  To customize the extensions or their weights, modify `fileExtensions` and `extensionWeights` in the script.

- **Directory Structure Parameters**: You can modify the `levels` map to change the folder/file count and margin settings for each complexity level:
  ```go
  var levels = map[int][3]int{
      1: {100, 500, 10},    // Low complexity: 100 folders, 500 files, 10% margin
      2: {1000, 5000, 10},  // Medium complexity: 1000 folders, 5000 files, 10% margin
      ...
  }
  ```

## Output

- The script will print out detailed metrics after generation:
  - **Directory Path**: The path where the structure was created.
  - **Folders and Files Created**: The number of folders and files generated, including the maximum limits and margin.
  - **Total Size**: The cumulative size of the generated files.
  - **Time Taken**: The total time taken to generate the structure.
  - **Operations per Second (OPS)**: A performance metric indicating the creation speed.

Example:
```
--- Generation Metrics ---
Directory structure generated in: /path/to/target_directory
Folders Created: 110 (Max Limit: 100, Margin: ±10%)
Files Created: 550 (Max Limit: 500, Margin: ±10%)
Total Size of Data Generated: 12000 bytes
Total Time Taken: 3.456 seconds
Operations per Second (OPS): 191.304
--------------------------
```

## Notes

- **Concurrency**: The script uses Go routines and synchronization to efficiently generate the structure in parallel.
- **Empty Directory Handling**: Any empty directory encountered will be filled with a placeholder file to ensure it is not left empty.
- **Randomness**: The script uses Go's `rand` package to ensure randomization of file names, extensions, and content.

## License

This script is released under [The Unlicense](https://unlicense.org/), placing it in the public domain and allowing free use without restrictions.
