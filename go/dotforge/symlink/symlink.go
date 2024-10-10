package symlink

import (
	"dotforge/config"
	"dotforge/utils"
	"fmt"
	"os"
	"path/filepath"
)

// CreateSymlinks creates symlinks based on the registered file types and ignored paths
func CreateSymlinks(source, target string) {
	// Resolve and normalize source directory
	absSource, err := utils.ResolveAndNormalizePath(source)
	if err != nil {
		fmt.Println("Error resolving source to absolute path:", err)
		return
	}

	// Resolve and normalize the target directory
	absTarget, err := utils.ResolveAndNormalizePath(target)
	if err != nil {
		fmt.Println("Error resolving target to absolute path:", err)
		return
	}

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
	filetypesMap := make(map[string]bool)
	for _, ft := range filetypes {
		filetypesMap[ft] = true
	}

	ignoredPathsMap := make(map[string]bool)
	for _, ip := range ignoredPaths {
		ignoredPathsMap[ip] = true
	}

	// Walk through source directory and create symlinks in target
	err = filepath.Walk(absSource, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Resolve the current path to its absolute form
		absPath, err := filepath.Abs(path)
		if err != nil {
			return err
		}

		// Skip ignored paths
		if ignoredPathsMap[absPath] {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Check if the file extension is in the approved list
		ext := filepath.Ext(info.Name())
		if filetypesMap[ext] {
			// Logic to create a symlink in the target directory
			targetPath := filepath.Join(absTarget, info.Name())
			err := os.Symlink(absPath, targetPath)
			if err != nil {
				fmt.Printf("Failed to create symlink for %s: %v\n", absPath, err)
			} else {
				fmt.Printf("Created symlink for %s -> %s\n", absPath, targetPath)
			}
		}

		return nil
	})

	// Error handling for the Walk function
	if err != nil {
		fmt.Println("Error creating symlinks:", err)
	}
}
