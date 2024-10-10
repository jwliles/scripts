package scanner

import (
	"dotforge/config"
	"dotforge/utils"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ScanDirectory scans the current directory and reports blocked or unaccepted files
func ScanDirectory(maxLines int) {
	var toBeSymlinked []string
	var blockedFolders []string
	var blockedFiles []string
	var unacceptedFiles []string

	// Read registered file types and ignored paths
	filetypes, err := utils.ReadLines(config.GetFiletypesFile())
	if err != nil {
		fmt.Println("Error reading file types:", err)
		return
	}

	ignoredPaths, err := utils.ReadLines(config.GetIgnoredPathsFile())
	if err != nil {
		fmt.Println("Error reading ignored paths:", err)
		return
	}

	// Create maps for faster lookup
	ignoredPathsMap := make(map[string]bool)
	for _, ip := range ignoredPaths {
		ignoredPathsMap[ip] = true
	}

	filetypesMap := make(map[string]bool)
	for _, ft := range filetypes {
		filetypesMap[ft] = true
	}

	// Walk through the current directory
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the current directory
		if path == "." {
			return nil
		}

		// Skip hidden files and directories
		if strings.HasPrefix(info.Name(), ".") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		absPath, err := filepath.Abs(path)
		if err != nil {
			return err
		}

		// Check if the path is in the ignored paths
		if ignoredPathsMap[absPath] {
			if info.IsDir() {
				blockedFolders = append(blockedFolders, path)
				return filepath.SkipDir // Skip the entire directory
			} else {
				blockedFiles = append(blockedFiles, path)
				return nil
			}
		}

		if info.IsDir() {
			return nil
		}

		// Check if the file extension matches any registered type
		ext := filepath.Ext(info.Name())
		if filetypesMap[ext] {
			toBeSymlinked = append(toBeSymlinked, path)
		} else {
			unacceptedFiles = append(unacceptedFiles, path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error scanning the directory:", err)
		return
	}

	// Display the results with a limit based on maxLines
	displayLines := func(header string, items []string) {
		fmt.Printf("\n%s:\n", header)
		for i, item := range items {
			if i >= maxLines {
				fmt.Printf("...and more. Showing first %d out of %d.\n", maxLines, len(items))
				break
			}
			fmt.Println("  -", item)
		}
	}

	displayLines("Files to be symlinked", toBeSymlinked)
	displayLines("Blocked folders", blockedFolders)
	displayLines("Blocked files", blockedFiles)
	displayLines("Files with unaccepted file types", unacceptedFiles)

	// Check if total lines exceed maxLines, open in editor if necessary
	totalLines := len(toBeSymlinked) + len(blockedFolders) + len(blockedFiles) + len(unacceptedFiles)
	if totalLines > maxLines {
		fmt.Printf("Results exceed %d lines, opening in editor...\n", maxLines)
		// Open results in a text editor here (implement as needed)
	}
}
