package main

import (
	"fmt"
	"os"
	modules "tetris-optimizer/modules"
)

func main() {
	content, err := modules.ProcessInput()
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(1)
	}

	// Parse tetrominoes
	tetrominoStrings, err := modules.ParseTetrominoes(content)
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(1)
	}

	// Validate tetrominoes format
	err = modules.ValidateTetrominoes(tetrominoStrings)
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(1)
	}

	// Create tetromino models
	var tetrominoes []*modules.Tetromino
	for _, tetrominoStr := range tetrominoStrings {
		tetromino := modules.NewTetromino(tetrominoStr)
		tetrominoes = append(tetrominoes, tetromino)
	}

	// For now, just print that we successfully created the models
	// This will be replaced with the solver in later tasks
	fmt.Printf("Successfully created %d tetromino models\n", len(tetrominoes))
	for i, t := range tetrominoes {
		fmt.Printf("Tetromino %d: %dx%d\n", i+1, t.Width(), t.Height())
	}
}
