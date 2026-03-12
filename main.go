package main

import (
	"fmt"
	"os"
	input "tetris-optimizer/modules"
)

func main() {
	content, err := input.ProcessInput()
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(1)
	}

	// Parse tetrominoes
	tetrominoes, err := input.ParseTetrominoes(content)
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(1)
	}

	// For now, just print that we successfully parsed the tetrominoes
	// This will be replaced with the full pipeline in later tasks
	fmt.Printf("Successfully parsed %d tetrominoes\n", len(tetrominoes))
}
