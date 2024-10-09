package output

import (
	"fmt"
	"go_readme/hashing"
	"go_readme/scanning"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// createReadmeFile creates a README.md file at the specified path
func createReadmeFile(path string) error {
	// Check if README.md already exists
	if _, err := os.Stat(path); err == nil {
		fmt.Printf("README.md already exists at %s\n", path)
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}

	// Create the README.md file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("# README\n")
	if err != nil {
		return err
	}

	return nil
}

// ParseTemplate parses the template by replacing placeholders with actual values
func ParseTemplate(templateContent, directory string, fileCount, dirCount int) string {
	// Get current date in ISO 8601 format
	currentDate := time.Now().Format("2006-01-02")

	// Extract base directory name (e.g., "notes")
	baseDir := filepath.Base(directory)

	// Replace placeholders in the template
	parsedContent := strings.ReplaceAll(templateContent, "{Directory}", baseDir)
	parsedContent = strings.ReplaceAll(parsedContent, "{?}", strconv.Itoa(dirCount)+" directories and "+strconv.Itoa(fileCount)+" files")
	parsedContent = strings.ReplaceAll(parsedContent, "{ISO 8601}", currentDate)

	return parsedContent
}

// GenerateAllReadmes processes all directories and generates `README.md` files as necessary
func GenerateAllReadmes(directoryFiles map[string][]string, subDirectories map[string][]string) error {
	// Create a map to mark which directories have been processed
	processedDirs := make(map[string]bool)

	for dir, files := range directoryFiles {
		subDirs := subDirectories[dir]
		// Generate or update README.md for the current directory
		if err := GenerateReadmeIfNecessary(dir, files, subDirs); err != nil {
			return err
		}

		processedDirs[dir] = true // Mark the current directory as processed

		// Process each subdirectory
		for _, subDir := range subDirs {
			subDirFiles := directoryFiles[subDir]
			subDirSubDirs := subDirectories[subDir]

			// Generate or update README.md for each subdirectory
			if err := GenerateReadmeIfNecessary(subDir, subDirFiles, subDirSubDirs); err != nil {
				return err
			}

			processedDirs[subDir] = true // Mark the current subdirectory as processed
		}
	}

	// Ensure that all directories have README.md files, including those without any files
	for dir := range subDirectories {
		if !processedDirs[dir] {
			readmePath := filepath.Join(dir, "README.md")
			if err := createReadmeFile(readmePath); err != nil {
				fmt.Printf("Error creating README.md in %s: %v\n", dir, err)
				continue
			}
			processedDirs[dir] = true // Mark these directories as processed
		}
	}
	return nil
}

// GenerateReadmeIfNecessary generates or updates the README.md for the given directory
func GenerateReadmeIfNecessary(directory string, files []string, subDirs []string) error {
	readmePath := filepath.Join(directory, "README.md")

	// Generate content for the potential README.md file
	generatedContent, err := GenerateReadmeContent(directory, files, subDirs)
	if err != nil {
		return err
	}

	// Debug: Print generated content for the README.md
	fmt.Printf("Generated content for README.md in %s:\n%s\n", directory, generatedContent)

	// Assume that a README.md needs to be created or updated
	needsUpdate := true

	// Condition 1: Check if README.md already exists and matches the generated content
	if _, err := os.Stat(readmePath); err == nil {
		existingContent, err := os.ReadFile(readmePath)
		if err != nil {
			return fmt.Errorf("could not read existing README.md: %v", err)
		}

		// Debug: Print existing content for the README.md
		fmt.Printf("Existing content for README.md in %s:\n%s\n", directory, string(existingContent))

		// If the existing content matches the generated content, no update is necessary
		if string(existingContent) == generatedContent {
			fmt.Printf("Skipping update for README.md in directory: %s (content unchanged)\n", directory)
			needsUpdate = false
		}
	}

	// Create or update README.md only if necessary
	if needsUpdate {
		err = os.WriteFile(readmePath, []byte(generatedContent), 0644)
		if err != nil {
			return fmt.Errorf("could not write to README.md: %v", err)
		}
		fmt.Printf("Generated/Updated README.md for directory: %s\n", directory)
	}

	return nil
}

// GenerateReadmeContent creates the content for a README.md file based on the current directory structure
func GenerateReadmeContent(directory string, files []string, subDirs []string) (string, error) {
	templatePath := "templates/readme_template.md"

	// Read the template content
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("could not read template: %v", err)
	}

	// Calculate file and directory counts (including this README file)
	fileCount := len(files) + 1 // Include this README file in the count
	dirCount := len(subDirs)

	// Parse the template content with actual data
	parsedContent := ParseTemplate(string(templateContent), directory, fileCount, dirCount)

	// Generate the file and subdirectory listings
	var fileLinks strings.Builder
	fileLinks.WriteString("\n## Directory Contents\n\n")

	// List files in the current directory
	for _, file := range files {
		relativePath := strings.TrimPrefix(file, directory+"/")
		fileLinks.WriteString(fmt.Sprintf("- [%s](%s)\n", filepath.Base(relativePath), relativePath))
	}

	// List subdirectories and their README files for navigation
	for _, subDir := range subDirs {
		subReadme := filepath.Join(subDir, "README.md")
		relativePath := strings.TrimPrefix(subReadme, directory+"/")
		fileLinks.WriteString(fmt.Sprintf("- [%s/README.md](%s)\n", filepath.Base(subDir), relativePath))
	}

	// Append the list of files and directories to the parsed content
	fullContent := parsedContent + fileLinks.String()

	// Compute content and path hashes
	contentHash := hashing.GenerateHash(fullContent)
	pathHash := hashing.HashFilePath(filepath.Join(directory, "README.md"))

	// Append hashes as markdown comments
	fullContent += fmt.Sprintf("\n<!-- file content hash:%s -->\n", contentHash)
	fullContent += fmt.Sprintf("<!-- file path hash:%s -->\n", pathHash)

	return fullContent, nil
}

// ProcessDirectories processes the directories and generates README.md files
func ProcessDirectories(rootDirectory string) error {
	// Scan the root directory recursively
	directoryFiles, subDirectories, err := scanning.ScanDirectory(rootDirectory)
	if err != nil {
		return fmt.Errorf("error scanning directory: %v", err)
	}

	// Generate README.md for each directory
	err = GenerateAllReadmes(directoryFiles, subDirectories)
	if err != nil {
		return fmt.Errorf("error generating README.md files: %v", err)
	}

	return nil
}
