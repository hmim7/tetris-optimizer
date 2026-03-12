package input

import (
	"strings"
)

// ValidateTetrominoes validates an array of tetromino blocks for format compliance
func ValidateTetrominoes(tetrominoes []string) error {
	if len(tetrominoes) == 0 {
		return &ValidationError{"no tetrominoes to validate"}
	}

	for i, tetromino := range tetrominoes {
		if err := ValidateTetromino(tetromino); err != nil {
			return &ValidationError{"tetromino " + string(rune(i+1)) + ": " + err.Error()}
		}
	}

	return nil
}

// ValidateTetromino validates a single tetromino block for format compliance
func ValidateTetromino(tetromino string) error {
	lines := strings.Split(tetromino, "\n")

	// Validate exactly 4 lines
	if len(lines) != 4 {
		return &ValidationError{"invalid block height"}
	}

	blockCount := 0
	var blocks [][]bool

	// Initialize 4x4 grid
	for i := 0; i < 4; i++ {
		blocks = append(blocks, make([]bool, 4))
	}

	// Validate each line and build block grid
	for row, line := range lines {
		// Validate line length is exactly 4
		if len(line) != 4 {
			return &ValidationError{"invalid line length"}
		}

		// Validate characters and count blocks
		for col, char := range line {
			if char != '.' && char != '#' {
				return &ValidationError{"invalid character"}
			}
			if char == '#' {
				blockCount++
				blocks[row][col] = true
			}
		}
	}

	// Validate exactly 4 blocks
	if blockCount != 4 {
		if blockCount < 4 {
			return &ValidationError{"too few blocks"}
		} else {
			return &ValidationError{"too many blocks"}
		}
	}

	// Validate connectivity
	if !isConnected(blocks) {
		return &ValidationError{"disconnected blocks"}
	}

	return nil
}

// ValidationError represents validation errors
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

// isConnected checks if all blocks in the 4x4 grid are connected edge-to-edge
func isConnected(blocks [][]bool) bool {
	// Find the first block to start flood fill
	startRow, startCol := -1, -1
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			if blocks[row][col] {
				startRow, startCol = row, col
				break
			}
		}
		if startRow != -1 {
			break
		}
	}

	// If no blocks found, return false
	if startRow == -1 {
		return false
	}

	// Create visited grid
	visited := make([][]bool, 4)
	for i := 0; i < 4; i++ {
		visited[i] = make([]bool, 4)
	}

	// Perform flood fill
	connectedCount := floodFill(blocks, visited, startRow, startCol)

	// All 4 blocks should be connected
	return connectedCount == 4
}

// floodFill performs recursive flood fill to count connected blocks
func floodFill(blocks [][]bool, visited [][]bool, row, col int) int {
	// Check bounds
	if row < 0 || row >= 4 || col < 0 || col >= 4 {
		return 0
	}

	// Check if already visited or not a block
	if visited[row][col] || !blocks[row][col] {
		return 0
	}

	// Mark as visited
	visited[row][col] = true
	count := 1

	// Check all 4 adjacent cells (up, down, left, right)
	count += floodFill(blocks, visited, row-1, col) // up
	count += floodFill(blocks, visited, row+1, col) // down
	count += floodFill(blocks, visited, row, col-1) // left
	count += floodFill(blocks, visited, row, col+1) // right

	return count
}
