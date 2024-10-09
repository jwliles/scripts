package scanning

import (
	"os"
	"path/filepath"
	"strings"
)

// ScanDirectory scans the provided root directory and returns files and subdirectories
func ScanDirectory(root string) (map[string][]string, map[string][]string, error) {
	directoryFiles := make(map[string][]string)
	subDirectories := make(map[string][]string)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip hidden directories and files (those starting with ".")
		if strings.HasPrefix(info.Name(), ".") {
			if info.IsDir() {
				return filepath.SkipDir // Skip the entire directory if it's hidden
			}
			return nil // Skip hidden files only
		}

		// Process directories and files
		if info.IsDir() {
			// Add detected subdirectory to its parent directory
			parentDir := filepath.Dir(path)
			if parentDir != path { // Avoid adding the root directory to itself
				subDirectories[parentDir] = append(subDirectories[parentDir], path)
			}
		} else {
			// Add detected file to its parent directory
			directoryFiles[filepath.Dir(path)] = append(directoryFiles[filepath.Dir(path)], path)
		}

		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return directoryFiles, subDirectories, nil
}
