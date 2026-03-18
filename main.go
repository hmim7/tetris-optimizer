package main

import (
	"fmt"
	"os"
	"tetris-optimizer/modules"
)

func main() {
	// 1. Strict Argument Validation
	// The subject requires "only one argument" (the file path).
	// os.Args[0] is the program name, os.Args[1] is the file.
	if len(os.Args) != 2 {
		fmt.Println("ERROR")
		os.Exit(1)
	}

	// 2. Read and Initial Process
	content, err := modules.ProcessInput()
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(1)
	}

	// 3. Parse tetrominoes
	tetrominoStrings, err := modules.ParseTetrominoes(content)
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(1)
	}

	// 4. Validate tetrominoes format
	err = modules.ValidateTetrominoes(tetrominoStrings)
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(1)
	}

	// 5. Create tetromino models
	var tetrominoes []*modules.Tetromino
	for _, tetrominoStr := range tetrominoStrings {
		tetromino := modules.NewTetromino(tetrominoStr)
		tetrominoes = append(tetrominoes, tetromino)
	}

	// 6. Solve the puzzle
	solution, err := modules.Solve(tetrominoes)
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(1)
	}

	// 7. Output the result
	// Ensure RenderOutput handles the required formatting
	fmt.Print(modules.RenderOutput(solution))
}
