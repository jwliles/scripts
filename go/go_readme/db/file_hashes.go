package db

import (
	"fmt"
)

// InsertFileHash logs a file's hash details to the file_hashes table
func InsertFileHash(dirPath, filePath, contentHash, pathHash string, lastModified int64) error {
	_, err := DB.Exec(
		"INSERT INTO file_hashes (dir_path, file_path, content_hash, path_hash, last_modified) VALUES (?, ?, ?, ?, ?)",
		dirPath, filePath, contentHash, pathHash, lastModified,
	)
	if err != nil {
		return fmt.Errorf("failed to insert file hash: %v", err)
	}
	return nil
}

// DeleteFileHashesInDir deletes all file hash records within a specified directory
func DeleteFileHashesInDir(dirPath string) error {
	pattern := fmt.Sprintf("%s%%", dirPath)
	_, err := DB.Exec("DELETE FROM file_hashes WHERE file_path LIKE ?", pattern)
	if err != nil {
		return fmt.Errorf("failed to delete file hashes in directory: %v", err)
	}
	return nil
}

// GetFileHashesInDir retrieves all file hashes directly under a specified directory
func GetFileHashesInDir(dirPath string) ([]struct {
	FilePath     string
	ContentHash  string
	PathHash     string
	LastModified int64
}, error) {
	pattern := fmt.Sprintf("%s%%", dirPath)

	rows, err := DB.Query(`
		SELECT file_path, content_hash, path_hash, last_modified
		FROM file_hashes
		WHERE file_path LIKE ?
	`, pattern)
	if err != nil {
		return nil, fmt.Errorf("failed to query file hashes: %v", err)
	}
	defer rows.Close()

	var results []struct {
		FilePath     string
		ContentHash  string
		PathHash     string
		LastModified int64
	}

	for rows.Next() {
		var result struct {
			FilePath     string
			ContentHash  string
			PathHash     string
			LastModified int64
		}
		if err := rows.Scan(&result.FilePath, &result.ContentHash, &result.PathHash, &result.LastModified); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		results = append(results, result)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return results, nil
}
