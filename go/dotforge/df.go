package main

import (
	"dotforge/cli"
	"fmt"
	"os"
	"strings"
)

// DefaultPathFile holds the path to the default configuration file
var defaultPathFile = os.Getenv("HOME") + "/.dotforge_config"

// ReadDefaultPath reads the default path from the configuration file
func ReadDefaultPath() string {
	data, err := os.ReadFile(defaultPathFile)
	if err != nil {
		return "~/dotforge-errors.md"
	}
	return strings.TrimSpace(string(data))
}

func main() {
	// Debug print to ensure we are in the main function
	fmt.Println("Starting Dotforge...")

	// Parse command-line flags and arguments
	cli.ParseFlags()
}

// getSourceAndTarget handles inference of source and target directories
func getSourceAndTarget(args []string) (string, string) {
	switch len(args) {
	case 0:
		defaultTarget := ReadDefaultPath()
		if defaultTarget == "" {
			fmt.Println("Error: No default target set and no target directory provided.")
			os.Exit(1) // Terminate the program
		}
		return ".", defaultTarget
	case 1:
		return ".", args[0]
	default:
		return args[0], args[1]
	}
}
