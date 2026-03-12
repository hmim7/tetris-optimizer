package main

import (
	"os"
	"testing"
	"tetris-optimizer/modules"
)

func TestValidateArgs(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantErr  bool
		expected string
	}{
		{
			name:    "No arguments",
			args:    []string{"program"},
			wantErr: true,
		},
		{
			name:    "Too many arguments",
			args:    []string{"program", "file1.txt", "file2.txt"},
			wantErr: true,
		},
		{
			name:     "Valid single argument",
			args:     []string{"program", "test.txt"},
			wantErr:  false,
			expected: "test.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original args
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			// Set test args
			os.Args = tt.args

			result, err := input.ValidateArgs()
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("ValidateArgs() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestReadFile(t *testing.T) {
	// Create a temporary test file
	testFile := "test_temp.txt"
	testContent := "test content"
	err := os.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(testFile)

	// Create an empty test file
	emptyFile := "empty_temp.txt"
	err = os.WriteFile(emptyFile, []byte(""), 0644)
	if err != nil {
		t.Fatalf("Failed to create empty test file: %v", err)
	}
	defer os.Remove(emptyFile)

	tests := []struct {
		name     string
		filename string
		wantErr  bool
		expected string
	}{
		{
			name:     "Valid file",
			filename: testFile,
			wantErr:  false,
			expected: testContent,
		},
		{
			name:     "Empty file",
			filename: emptyFile,
			wantErr:  true,
		},
		{
			name:     "Non-existent file",
			filename: "nonexistent.txt",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := input.ReadFile(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("ReadFile() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestProcessInput(t *testing.T) {
	// Create a temporary test file
	testFile := "test_process.txt"
	testContent := "process test content"
	err := os.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(testFile)

	tests := []struct {
		name     string
		args     []string
		wantErr  bool
		expected string
	}{
		{
			name:    "No arguments",
			args:    []string{"program"},
			wantErr: true,
		},
		{
			name:    "Too many arguments",
			args:    []string{"program", "file1.txt", "file2.txt"},
			wantErr: true,
		},
		{
			name:     "Valid file processing",
			args:     []string{"program", testFile},
			wantErr:  false,
			expected: testContent,
		},
		{
			name:    "Non-existent file",
			args:    []string{"program", "nonexistent.txt"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original args
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			// Set test args
			os.Args = tt.args

			result, err := input.ProcessInput()
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("ProcessInput() = %v, want %v", result, tt.expected)
			}
		})
	}
}
