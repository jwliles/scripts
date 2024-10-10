package utils

import (
	"fmt"
	"path/filepath"
)

// ResolveAndNormalizePath resolves a given path to its absolute path
// and normalizes it to ensure consistent comparison
func ResolveAndNormalizePath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("error resolving path to absolute: %w", err)
	}
	return filepath.Clean(absPath), nil
}
