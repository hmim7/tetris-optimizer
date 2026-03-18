package modules

import (
	"fmt"
	"math"
)

// solvePiece represents a single tetromino with
// its original input order and assigned letter for the backtracking solver.
type solvePiece struct {
	tetromino *Tetromino
	index     int
	letter    byte
}

// Solve calculates the smallest square board that can contain all tetrominoes.
// It uses an iterative deepening approach, increasing the square size until a solution is found.
func Solve(tetrominoes []*Tetromino) ([]string, error) {
	if len(tetrominoes) == 0 {
		return nil, fmt.Errorf("no tetrominoes")
	}
	if len(tetrominoes) > 26 {
		return nil, fmt.Errorf("too many tetrominoes")
	}

	// Wrap tetrominoes in solvePiece to track their original order and assign letters.
	pieces := make([]solvePiece, 0, len(tetrominoes))
	for i, t := range tetrominoes {
		pieces = append(pieces, solvePiece{
			tetromino: t,
			index:     i,
			letter:    byte('A' + i),
		})
	}

	// Place pieces deterministically in input order (A, B, C...).

	// Start searching from the absolute mathematical minimum size possible.
	minSize := minimalSquareSize(len(pieces) * 4)
	for size := minSize; ; size++ {
		board := make([]byte, size*size)
		for i := range board {
			board[i] = '.'
		}

		if backtrackPlace(board, size, pieces, 0) {
			return boardToStrings(board, size), nil
		}
	}
}

// minimalSquareSize returns the smallest integer side length needed to hold the total block count.
func minimalSquareSize(blocks int) int {
	return int(math.Ceil(math.Sqrt(float64(blocks))))
}

// backtrackPlace is a recursive function that attempts to place pieces one by one.
// It uses depth-first search to find a valid board configuration.
func backtrackPlace(board []byte, size int, pieces []solvePiece, pieceIndex int) bool {
	// Base case: all pieces have been successfully placed on the board.
	if pieceIndex == len(pieces) {
		return true
	}

	// For larger boards, choose the most constrained piece among the remaining ones.
	// This improves performance without affecting small-board deterministic outputs.
	bestIdx := -1
	swapped := false
	if size >= 7 {
		bestCount := math.MaxInt32
		for i := pieceIndex; i < len(pieces); i++ {
			count := possiblePlacementCount(board, size, pieces[i].tetromino)
			if count == 0 {
				return false
			}
			if bestIdx == -1 || count < bestCount || (count == bestCount && pieces[i].index < pieces[bestIdx].index) {
				bestCount = count
				bestIdx = i
			}
		}
		if bestIdx != -1 && bestIdx != pieceIndex {
			pieces[pieceIndex], pieces[bestIdx] = pieces[bestIdx], pieces[pieceIndex]
			swapped = true
		}
	}

	piece := pieces[pieceIndex]
	t := piece.tetromino

	// Pre-calculate boundary limits to avoid checking coordinates outside the board.
	maxRow := size - t.Height()
	maxCol := size - t.Width()

	// Try every valid top-left coordinate for the current piece.
	for row := 0; row <= maxRow; row++ {
		for col := 0; col <= maxCol; col++ {
			if !canPlaceAt(board, size, t, row, col) {
				continue
			}

			// Place the piece, move to the next, and undo if the path leads to a dead end.
			placeAt(board, size, t, row, col, piece.letter)
			if backtrackPlace(board, size, pieces, pieceIndex+1) {
				return true
			}
			removeAt(board, size, t, row, col)
		}
	}

	if swapped {
		pieces[pieceIndex], pieces[bestIdx] = pieces[bestIdx], pieces[pieceIndex]
	}
	return false
}

// possiblePlacementCount counts how many positions a tetromino can be placed at
// on the current board state.
func possiblePlacementCount(board []byte, size int, t *Tetromino) int {
	count := 0
	maxRow := size - t.Height()
	maxCol := size - t.Width()
	for row := 0; row <= maxRow; row++ {
		for col := 0; col <= maxCol; col++ {
			if canPlaceAt(board, size, t, row, col) {
				count++
			}
		}
	}
	return count
}

// canPlaceAt checks if a tetromino's blocks overlap with existing pieces or the board edges.
func canPlaceAt(board []byte, size int, t *Tetromino, startRow, startCol int) bool {
	for _, block := range t.Blocks {
		r := startRow + block.Row
		c := startCol + block.Col
		// Verify coordinates are within board bounds and target cell is empty.
		if r < 0 || r >= size || c < 0 || c >= size {
			return false
		}
		if board[r*size+c] != '.' {
			return false
		}
	}
	return true
}

// placeAt writes the piece's letter to the board for each of its four blocks.
func placeAt(board []byte, size int, t *Tetromino, startRow, startCol int, letter byte) {
	for _, block := range t.Blocks {
		r := startRow + block.Row
		c := startCol + block.Col
		board[r*size+c] = letter
	}
}

// removeAt clears the piece's blocks from the board, resetting them to '.' characters.
func removeAt(board []byte, size int, t *Tetromino, startRow, startCol int) {
	for _, block := range t.Blocks {
		r := startRow + block.Row
		c := startCol + block.Col
		board[r*size+c] = '.'
	}
}

// boardToStrings converts the flat 1D byte board into a 2D-like slice of strings for final output.
func boardToStrings(board []byte, size int) []string {
	out := make([]string, size)
	for r := 0; r < size; r++ {
		out[r] = string(board[r*size : (r+1)*size])
	}
	return out
}
