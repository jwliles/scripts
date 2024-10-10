package config

import (
	"dotforge/utils"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	defaultPathFile  = os.Getenv("HOME") + "/.dotforge_config"
	filetypesFile    = os.Getenv("HOME") + "/.dotforge_filetypes"
	ignoredPathsFile = os.Getenv("HOME") + "/.dotforge_ignored_paths"
	maxLines         int
)

// GetFiletypesFile Getters for File Paths
func GetFiletypesFile() string {
	return filetypesFile
}

func GetIgnoredPathsFile() string {
	return ignoredPathsFile
}

// SetDefaultPath Set the default path
func SetDefaultPath(path string) error {
	return os.WriteFile(defaultPathFile, []byte(path), 0644)
}

func ReadDefaultPath() string {
	data, err := os.ReadFile(defaultPathFile)
	if err != nil {
		return "~/dotforge-errors.md"
	}
	return strings.TrimSpace(string(data))
}

// AddFiletypes Add File Types
func AddFiletypes(extensions []string) {
	for _, extension := range extensions {
		extension = strings.TrimSpace(extension)
		if extension == "" {
			continue
		}
		if fileExists(GetFiletypesFile(), extension) {
			fmt.Printf("File type '%s' is already approved.\n", extension)
		} else {
			appendToFile(GetFiletypesFile(), extension)
			fmt.Printf("File type '%s' added to the approved list.\n", extension)
		}
	}
}

// RemoveFiletypes Remove File Types
func RemoveFiletypes(extensions []string) {
	for _, extension := range extensions {
		RemoveItemFromList(GetFiletypesFile(), extension)
	}
}

// ListFiletypes List all approved file types
func ListFiletypes() {
	ListItems(GetFiletypesFile(), "Approved File Types")
}

// AddIgnoredPaths Add Ignored Paths
func AddIgnoredPaths(paths []string) {
	for _, p := range paths {
		absPath, err := utils.ResolveAndNormalizePath(p)
		if err != nil {
			fmt.Printf("Error resolving path to absolute: %s: %v\n", p, err)
			continue
		}
		if fileExists(GetIgnoredPathsFile(), absPath) {
			fmt.Printf("Path '%s' is already blocked.\n", absPath)
		} else {
			appendToFile(GetIgnoredPathsFile(), absPath)
			fmt.Printf("Path '%s' added to the blocked list.\n", absPath)
		}
	}
}

// RemoveIgnoredPaths Remove Ignored Paths
func RemoveIgnoredPaths(paths []string) {
	for _, path := range paths {
		RemoveItemFromList(GetIgnoredPathsFile(), path)
	}
}

// ListIgnoredPaths List all ignored paths
func ListIgnoredPaths() {
	ListItems(GetIgnoredPathsFile(), "Blocked Paths")
}

// ListItems List items in a file
func ListItems(filePath, header string) {
	lines, err := readLines(filePath)
	if err != nil {
		fmt.Println("Error reading the list:", err)
		return
	}

	fmt.Printf("\n%s:\n", header)
	if len(lines) == 0 {
		fmt.Println("  No items found.")
	} else {
		for _, line := range lines {
			fmt.Println("  -", line)
		}
	}
}

// RemoveItemFromList Remove a specific item from a list
func RemoveItemFromList(filePath, item string) {
	item = strings.TrimSpace(item)
	if item == "" {
		fmt.Println("No item specified to remove.")
		return
	}

	lines, err := readLines(filePath)
	if err != nil {
		fmt.Println("Error reading the list:", err)
		return
	}

	var updatedLines []string
	found := false
	for _, line := range lines {
		if strings.TrimSpace(line) != item {
			updatedLines = append(updatedLines, line)
		} else {
			found = true
		}
	}

	if !found {
		fmt.Printf("Item '%s' not found in the list.\n", item)
		return
	}

	err = os.WriteFile(filePath, []byte(strings.Join(updatedLines, "\n")+"\n"), 0644)
	if err != nil {
		fmt.Println("Error updating the list:", err)
		return
	}

	fmt.Printf("Item '%s' removed successfully.\n", item)
}

// Utility Functions for File Handling
func fileExists(filePath, item string) bool {
	lines, err := readLines(filePath)
	if err != nil {
		return false
	}
	for _, line := range lines {
		if strings.TrimSpace(line) == item {
			return true
		}
	}
	return false
}

func appendToFile(filePath, text string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	if _, err = file.WriteString(text + "\n"); err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func readLines(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	return lines, nil
}

func RemoveOldSymlinks(targetDir string) {
	err := filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Mode()&os.ModeSymlink != 0 {
			if err := os.Remove(path); err != nil {
				fmt.Printf("Failed to remove symlink: %s, error: %v\n", path, err)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error removing old symlinks:", err)
	}
}
