package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the SQLite database connection and tables
func InitDB(dbPath string) {
	var err error
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	log.Printf("Database initialized at path: %s\n", dbPath)

	createTables()
}

// createTables creates necessary tables in the database if they don't exist
func createTables() {
	createFileHashesTable := `
	CREATE TABLE IF NOT EXISTS file_hashes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		dir_path TEXT NOT NULL,
		file_path TEXT NOT NULL,
		content_hash TEXT NOT NULL,
		path_hash TEXT NOT NULL,
		last_modified INTEGER
	);`

	createMetricsTable := `
	CREATE TABLE IF NOT EXISTS metrics (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		metric_name TEXT NOT NULL,
		metric_value REAL NOT NULL,
		timestamp INTEGER NOT NULL
	);`

	if _, err := DB.Exec(createFileHashesTable); err != nil {
		log.Fatalf("Failed to create 'file_hashes' table: %v", err)
	}
	log.Println("Table 'file_hashes' created or verified.")

	if _, err := DB.Exec(createMetricsTable); err != nil {
		log.Fatalf("Failed to create 'metrics' table: %v", err)
	}
	log.Println("Table 'metrics' created or verified.")
}

// CloseDB closes the SQLite database connection
func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatalf("Failed to close database: %v", err)
	}
}
