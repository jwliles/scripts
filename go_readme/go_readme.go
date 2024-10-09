package main

import (
	"flag"
	"fmt"
	"go_readme/db"
	"go_readme/output"
	"log"
	"os"
)

func main() {
	root, dbPath := parseFlags()

	fmt.Printf("Starting scan on path: %s\n", *root)

	initializeDB(*dbPath)
	defer db.CloseDB()

	processDirectories(*root)
}

func parseFlags() (root *string, dbPath *string) {
	root = flag.String("path", ".", "The root directory to scan.")
	dbPath = flag.String("db", "readme_metrics.db", "The path to the SQLite database file.")
	flag.Parse()
	return
}

func initializeDB(dbPath string) {
	db.InitDB(dbPath)
}

func processDirectories(root string) {
	// Check if the directory exists and is accessible
	if _, err := os.Stat(root); os.IsNotExist(err) {
		log.Fatalf("The directory %s does not exist: %v", root, err)
	} else if err != nil {
		log.Fatalf("Failed to access the directory %s: %v", root, err)
	}

	// Call the ProcessDirectories function and handle errors
	err := output.ProcessDirectories(root)
	if err != nil {
		log.Fatalf("Error processing directories: %v", err)
	}
}
