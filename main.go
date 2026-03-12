package main

import (
	"fmt"
	"os"
	"tetris-optimizer/modules"
)

func main() {
	content, err := input.ProcessInput()
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(1)
	}

	// For now, just print that we successfully read the file
	// This will be replaced with the full pipeline in later tasks
	fmt.Printf("Successfully read file with %d characters\n", len(content))
}
