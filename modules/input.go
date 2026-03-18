package modules

import (
	"fmt"
	"os"
	"strings"
)

// ProcessInput handles the complete input validation and reading pipeline
func ProcessInput() (string, error) {
	filename, err := ValidateArgs()
	if err != nil {
		return "", err
	}

	content, err := ReadFile(filename)
	if err != nil {
		return "", err
	}

	return content, nil
}

// ValidateArgs validates command line arguments and returns the filename
func ValidateArgs() (string, error) {
	if len(os.Args) != 2 {
		return "", fmt.Errorf("invalid argument count")
	}
	return os.Args[1], nil
}

// ReadFile reads and validates the input file
func ReadFile(filename string) (string, error) {
	// Check if file exists and get file info
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return "", fmt.Errorf("file access error")
	}

	// Check if it's a directory
	if fileInfo.IsDir() {
		return "", fmt.Errorf("path is directory")
	}

	// Check if file is empty
	if fileInfo.Size() == 0 {
		return "", fmt.Errorf("empty file")
	}

	// Read file content
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("file read error")
	}

	// Check for null bytes (binary file detection)
	for _, b := range content {
		if b == 0 {
			return "", fmt.Errorf("binary file detected")
		}
	}

	// Convert to string and validate content
	contentStr := string(content)

	// Check for valid characters only (., #, newline)
	for _, char := range contentStr {
		if char != '.' && char != '#' && char != '\n' && char != '\r' {
			return "", fmt.Errorf("invalid characters in file")
		}
	}

	// Normalize line endings (convert \r\n to \n)
	contentStr = strings.ReplaceAll(contentStr, "\r\n", "\n")
	contentStr = strings.ReplaceAll(contentStr, "\r", "\n")

	return contentStr, nil
}
