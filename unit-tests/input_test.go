package unittests

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

			result, err := modules.ValidateArgs()
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
	// Create a temporary test file with valid content
	testFile := "test_temp.txt"
	testContent := "...#\n...#\n...#\n...#"
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

	// Create a binary file with null bytes
	binaryFile := "binary_temp.txt"
	binaryContent := []byte{0x00, 0x01, 0x02, 0x03}
	err = os.WriteFile(binaryFile, binaryContent, 0644)
	if err != nil {
		t.Fatalf("Failed to create binary test file: %v", err)
	}
	defer os.Remove(binaryFile)

	// Create a file with invalid characters
	invalidFile := "invalid_temp.txt"
	invalidContent := "...#\n@@@#\n...#\n...#"
	err = os.WriteFile(invalidFile, []byte(invalidContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create invalid test file: %v", err)
	}
	defer os.Remove(invalidFile)

	// Create a file with Windows line endings
	windowsFile := "windows_temp.txt"
	windowsContent := "...#\r\n...#\r\n...#\r\n...#"
	err = os.WriteFile(windowsFile, []byte(windowsContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create Windows test file: %v", err)
	}
	defer os.Remove(windowsFile)

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
		{
			name:     "Binary file with null bytes",
			filename: binaryFile,
			wantErr:  true,
		},
		{
			name:     "File with invalid characters",
			filename: invalidFile,
			wantErr:  true,
		},
		{
			name:     "Windows line endings (normalized)",
			filename: windowsFile,
			wantErr:  false,
			expected: "...#\n...#\n...#\n...#",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := modules.ReadFile(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("ReadFile() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestProcessInput(t *testing.T) {
	// Create a temporary test file with valid tetromino content
	testFile := "test_process.txt"
	testContent := "...#\n...#\n...#\n...#"
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

			result, err := modules.ProcessInput()
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("ProcessInput() = %q, want %q", result, tt.expected)
			}
		})
	}
}
