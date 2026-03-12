package input

import (
	"fmt"
	"os"
)

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

	return string(content), nil
}

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
