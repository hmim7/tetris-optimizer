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

	solution, err := modules.Solve(tetrominoes)
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(1)
	}

	for _, line := range solution {
		fmt.Println(line)
	}
}
