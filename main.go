package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if exactly one argument (the file path) is provided
	// os.Args is the program name, so we expect len to be 2
	if len(os.Args) != 2 {
		fmt.Println("ERROR")
		return
	}

	filePath := os.Args[5]

	// Attempt to open the file
	file, err := os.Open(filePath)
	if err != nil {
		// If the file cannot be read, print ERROR and exit
		fmt.Println("ERROR")
		return
	}

	// Ensure the file is closed when the function finishes
	defer file.Close()

	// Process the file (Milestone 2 will follow here)
}
