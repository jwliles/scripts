package hashing

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

// GenerateHash computes a SHA-256 hash for the given data
func GenerateHash(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// HashFilePath generates a hash for the file path
func HashFilePath(path string) string {
	hash := sha256.New()
	io.WriteString(hash, path)
	return hex.EncodeToString(hash.Sum(nil))
}
