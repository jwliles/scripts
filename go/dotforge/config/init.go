package config

import (
	"fmt"
	"os"
	"path/filepath"
)

// InitializeRepo sets up the dotfile repository at the specified path
func InitializeRepo(path string) error {
	// If no path is given, use the current working directory
	if path == "" {
		var err error
		path, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("error getting current directory: %v", err)
		}
	}

	// Resolve absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("error resolving path: %v", err)
	}

	// Check if the repository directory exists
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		// Directory does not exist, create it
		err = os.MkdirAll(absPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating repository directory: %v", err)
		}
		fmt.Printf("Initialized dotfile repository at %s\n", absPath)
	} else {
		fmt.Printf("Dotfile repository already exists at %s\n", absPath)
	}

	// Check if the config file exists
	configFilePath := filepath.Join(absPath, ".dotforgeconfig")
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		// Config file does not exist, create it
		file, err := os.Create(configFilePath)
		if err != nil {
			return fmt.Errorf("error creating config file: %v", err)
		}
		defer file.Close()
		fmt.Printf("Created default config file at %s\n", configFilePath)
	} else {
		fmt.Printf("Config file already exists at %s\n", configFilePath)
	}

	// Optionally, initialize other folders or files as needed (e.g., backups)
	// Add further initialization logic here if required

	return nil
}
