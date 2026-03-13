package modules

import (
	"strings"
)

// Coordinate represents a 2D position
type Coordinate struct {
	Row, Col int
}

// Tetromino represents a tetromino shape as a collection of coordinates
type Tetromino struct {
	Blocks []Coordinate
	width  int
	height int
}

// NewTetromino creates a new tetromino from a 4x4 string representation
func NewTetromino(tetrominoStr string) *Tetromino {
	lines := strings.Split(tetrominoStr, "\n")
	var blocks []Coordinate

	// Extract block coordinates
	for row, line := range lines {
		for col, char := range line {
			if char == '#' {
				blocks = append(blocks, Coordinate{Row: row, Col: col})
			}
		}
	}

	// Create tetromino and normalize
	tetromino := &Tetromino{Blocks: blocks}
	tetromino.normalize()
	return tetromino
}

// normalize adjusts coordinates so the top-left block is at origin (0,0)
func (t *Tetromino) normalize() {
	if len(t.Blocks) == 0 {
		return
	}

	// Find minimum row and column
	minRow, minCol := t.Blocks[0].Row, t.Blocks[0].Col
	for _, block := range t.Blocks {
		if block.Row < minRow {
			minRow = block.Row
		}
		if block.Col < minCol {
			minCol = block.Col
		}
	}

	// Adjust all coordinates
	for i := range t.Blocks {
		t.Blocks[i].Row -= minRow
		t.Blocks[i].Col -= minCol
	}

	// Calculate dimensions
	t.calculateDimensions()
}

// calculateDimensions computes width and height of the tetromino
func (t *Tetromino) calculateDimensions() {
	if len(t.Blocks) == 0 {
		t.width, t.height = 0, 0
		return
	}

	maxRow, maxCol := 0, 0
	for _, block := range t.Blocks {
		if block.Row > maxRow {
			maxRow = block.Row
		}
		if block.Col > maxCol {
			maxCol = block.Col
		}
	}

	t.width = maxCol + 1
	t.height = maxRow + 1
}

// Width returns the width of the tetromino
func (t *Tetromino) Width() int {
	return t.width
}

// Height returns the height of the tetromino
func (t *Tetromino) Height() int {
	return t.height
}

// CanPlaceAt checks if the tetromino can be placed at the given position on a board
func (t *Tetromino) CanPlaceAt(boardSize, startRow, startCol int) bool {
	for _, block := range t.Blocks {
		newRow := startRow + block.Row
		newCol := startCol + block.Col

		// Check bounds
		if newRow < 0 || newRow >= boardSize || newCol < 0 || newCol >= boardSize {
			return false
		}
	}
	return true
}

// GetAbsoluteCoordinates returns the absolute coordinates when placed at a position
func (t *Tetromino) GetAbsoluteCoordinates(startRow, startCol int) []Coordinate {
	var coords []Coordinate
	for _, block := range t.Blocks {
		coords = append(coords, Coordinate{
			Row: startRow + block.Row,
			Col: startCol + block.Col,
		})
	}
	return coords
}
