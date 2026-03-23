package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"testing"
	"time"
)

var testBinaryPath string

func TestMain(m *testing.M) {
	// Create a temporary file for the built executable
	tmpFile, err := os.CreateTemp("", "tetris-optimizer-test-*.exe")
	if err != nil {
		fmt.Println("failed to create temp file for test binary:", err)
		os.Exit(1)
	}
	tmpFile.Close()
	testBinaryPath = tmpFile.Name()

	// Build the binary once
	cmd := exec.Command("go", "build", "-o", testBinaryPath, ".")
	if err := cmd.Run(); err != nil {
		fmt.Println("failed to build test binary:", err)
		os.Remove(testBinaryPath)
		os.Exit(1)
	}

	// Run all tests
	exitCode := m.Run()

	// Clean up the binary after tests are done
	os.Remove(testBinaryPath)
	os.Exit(exitCode)
}

// TestMain_Success tests the happy path of the main function.
// By passing a known valid file, it executes the entire pipeline
// without hitting os.Exit(1), giving us excellent code coverage.
func TestMain_Success(t *testing.T) {
	// 1. Save original os.Args and restore them when the test finishes
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// 2. Mock the command line arguments
	os.Args = []string{"tetris-optimizer", "samples/goodexample00.txt"}

	// 3. Hijack os.Stdout to capture the output of main()
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// 4. Run the program
	main()

	// 5. Close the pipe and restore normal standard output
	w.Close()
	os.Stdout = oldStdout

	// 6. Read what was captured
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// 7. Verify the result (goodexample00.txt should be a 2x2 square of 'A')
	expected := "AA\nAA\n"
	if output != expected {
		t.Errorf("Expected output:\n%q\nGot:\n%q", expected, output)
	}
}

// TestMain_ErrorPaths tests the scenarios that trigger os.Exit(1)
// by calling the pre-built executable directly.
func TestMain_ErrorPaths(t *testing.T) {
	// These are the bad argument scenarios we want to test
	tests := []struct {
		name string
		args []string
	}{
		{"No Arguments", []string{}},
		{"Too Many Arguments", []string{"file1.txt", "file2.txt"}},
		{"File Does Not Exist", []string{"ghost_file.txt"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute the pre-built binary
			cmd := exec.Command(testBinaryPath, tt.args...)
			output, err := cmd.CombinedOutput()

			// We EXPECT an error because os.Exit(1) tells the OS the program failed
			if e, ok := err.(*exec.ExitError); ok && !e.Success() {
				if !bytes.Contains(output, []byte("ERROR")) {
					t.Errorf("Expected output to contain 'ERROR', got: %q", string(output))
				}
			} else {
				t.Fatalf("Process ran successfully but we expected it to fail! err: %v", err)
			}
		})
	}
}

// TestMain_HardExam tests the performance and correctness of the 12-piece hard exam.
func TestMain_HardExam(t *testing.T) {
	start := time.Now()

	cmd := exec.Command(testBinaryPath, "samples/hardexam.txt")
	output, err := cmd.CombinedOutput()

	duration := time.Since(start)

	if err != nil {
		t.Fatalf("Expected successful execution, got error: %v, output: %s", err, string(output))
	}

	// The PRD specifies hardexam.txt should solve in <= 5 seconds.
	if duration.Seconds() > 5 {
		t.Errorf("Performance failed: execution took %v, which exceeds the 5-second limit", duration)
	}

	// The expected output should contain exactly 1 empty space ('.')
	dotCount := bytes.Count(output, []byte("."))
	if dotCount != 1 {
		t.Errorf("Expected exactly 1 empty space ('.'), got %d", dotCount)
	}
}

// TestMain_MaxPieces tests the absolute maximum limit of the program (26 tetrominoes).
func TestMain_MaxPieces(t *testing.T) {
	// This allows you to skip this 8+ second test during rapid development
	// by running: go test -short ./...
	if testing.Short() {
		t.Skip("Skipping 26-piece integration test in short mode")
	}

	start := time.Now()
	cmd := exec.Command(testBinaryPath, "samples/goodexample14.txt")
	output, err := cmd.CombinedOutput()
	duration := time.Since(start)

	if err != nil {
		t.Fatalf("Expected successful execution, got error: %v, output: %s", err, string(output))
	}

	// While not strictly defined in PRD, we ensure it doesn't take an eternity.
	if duration.Seconds() > 15 {
		t.Errorf("Performance warning: execution took %v, ideally should be < 15s", duration)
	}

	// 26 pieces = 104 blocks. Minimum square is 11x11 = 121 cells.
	// Expected empty spaces: 121 - 104 = 17 dots.
	dotCount := bytes.Count(output, []byte("."))
	if dotCount != 17 {
		t.Errorf("Expected exactly 17 empty spaces ('.'), got %d", dotCount)
	}
}
