package modules

import (
	"fmt"
	"math"
)

// solvePiece represents a single tetromino with
// its original input order and assigned letter for the backtracking solver.
type solvePiece struct {
	tetromino   *Tetromino
	index       int
	letter      byte
	mask        [4]uint16 // bitmask representation for fast placement
	width       int
	height      int
	identicalTo int // Tracks if this shape is a duplicate to prune mirror states
	placedPos   int
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
		var mask [4]uint16
		for _, b := range t.Blocks {
			mask[b.Row] |= (1 << b.Col)
		}

		// Pruning: Find if this piece is exactly identical to a previous one
		identicalTo := -1
		for j := i - 1; j >= 0; j-- {
			if pieces[j].width == t.Width() && pieces[j].height == t.Height() && pieces[j].mask == mask {
				identicalTo = j
				break
			}
		}

		pieces = append(pieces, solvePiece{
			tetromino:   t,
			index:       i,
			letter:      byte('A' + i),
			mask:        mask,
			width:       t.Width(),
			height:      t.Height(),
			identicalTo: identicalTo,
		})
	}

	minSize := int(math.Ceil(math.Sqrt(float64(len(pieces) * 4))))
	for size := minSize; ; size++ {
		var boardMask [20]uint16 // 20 allows bounds safety without checks
		board := make([]byte, size*size)
		for i := range board {
			board[i] = '.'
		}

		freeSpaces := size*size - len(pieces)*4
		if backtrackPlace(board, &boardMask, size, pieces, 0, freeSpaces) {
			return boardToStrings(board, size), nil
		}
	}
}

// backtrackPlace is a recursive function that attempts to place pieces one by one.
// It uses depth-first search to find a valid board configuration.
func backtrackPlace(board []byte, boardMask *[20]uint16, size int, pieces []solvePiece, pieceIndex int, freeSpaces int) bool {
	// Base case: all pieces have been successfully placed on the board.
	if pieceIndex == len(pieces) {
		return true
	}

	// --- PRUNING ---
	// Flood fill to check if the remaining empty spaces can mathematically hold the remaining pieces.
	if !isValidBoardStateMask(boardMask, size, freeSpaces) {
		return false
	}

	p := &pieces[pieceIndex]
	maxRow := size - p.height
	maxCol := size - p.width

	startRow := 0
	startCol := 0
	// If identical to a previous piece, force order to avoid exploring duplicate symmetrical states.
	if p.identicalTo != -1 {
		prevPos := pieces[p.identicalTo].placedPos
		startRow = prevPos / size
		startCol = prevPos % size
	}

	// Try every valid top-left coordinate for the current piece.
	for row := startRow; row <= maxRow; row++ {
		cStart := 0
		if row == startRow {
			cStart = startCol
		}
		for col := cStart; col <= maxCol; col++ {
			// Fast bitmask collision check.
			if (boardMask[row]&(p.mask[0]<<col)) != 0 ||
				(boardMask[row+1]&(p.mask[1]<<col)) != 0 ||
				(boardMask[row+2]&(p.mask[2]<<col)) != 0 ||
				(boardMask[row+3]&(p.mask[3]<<col)) != 0 {
				continue
			}

			// Place piece in bitmask
			boardMask[row] |= (p.mask[0] << col)
			boardMask[row+1] |= (p.mask[1] << col)
			boardMask[row+2] |= (p.mask[2] << col)
			boardMask[row+3] |= (p.mask[3] << col)

			// Place piece in byte board
			for _, b := range p.tetromino.Blocks {
				board[(row+b.Row)*size+(col+b.Col)] = p.letter
			}

			p.placedPos = row*size + col

			if backtrackPlace(board, boardMask, size, pieces, pieceIndex+1, freeSpaces) {
				return true
			}

			// Remove piece from bitmask
			boardMask[row] &= ^(p.mask[0] << col)
			boardMask[row+1] &= ^(p.mask[1] << col)
			boardMask[row+2] &= ^(p.mask[2] << col)
			boardMask[row+3] &= ^(p.mask[3] << col)

			// Remove piece from byte board
			for _, b := range p.tetromino.Blocks {
				board[(row+b.Row)*size+(col+b.Col)] = '.'
			}
		}
	}

	return false
}

// boardToStrings converts the flat 1D byte board into a 2D-like slice of strings for final output.
func boardToStrings(board []byte, size int) []string {
	out := make([]string, size)
	for r := 0; r < size; r++ {
		out[r] = string(board[r*size : (r+1)*size])
	}
	return out
}

// isValidBoardStateMask uses flood fill on the bitmask to find empty connected components.
// Since every tetromino occupies exactly 4 cells, an empty component
// of size S will leave at least (S % 4) cells unfillable.
// If the total unfillable cells exceed the number of free spaces, the board is invalid.
func isValidBoardStateMask(boardMask *[20]uint16, size, freeSpaces int) bool {
	var visited [20]uint16
	unfillable := 0

	var stack [512]uint16

	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			if (boardMask[r]&(1<<c)) == 0 && (visited[r]&(1<<c)) == 0 {
				compSize := 0
				top := 1
				stack[0] = uint16(r<<8 | c)
				visited[r] |= (1 << c)

				for top > 0 {
					top--
					curr := stack[top]
					cr := int(curr >> 8)
					cc := int(curr & 255)
					compSize++

					// Up
					if cr > 0 && (boardMask[cr-1]&(1<<cc)) == 0 && (visited[cr-1]&(1<<cc)) == 0 {
						visited[cr-1] |= (1 << cc)
						stack[top] = uint16((cr-1)<<8 | cc)
						top++
					}
					// Down
					if cr < size-1 && (boardMask[cr+1]&(1<<cc)) == 0 && (visited[cr+1]&(1<<cc)) == 0 {
						visited[cr+1] |= (1 << cc)
						stack[top] = uint16((cr+1)<<8 | cc)
						top++
					}
					// Left
					if cc > 0 && (boardMask[cr]&(1<<(cc-1))) == 0 && (visited[cr]&(1<<(cc-1))) == 0 {
						visited[cr] |= (1 << (cc - 1))
						stack[top] = uint16(cr<<8 | (cc - 1))
						top++
					}
					// Right
					if cc < size-1 && (boardMask[cr]&(1<<(cc+1))) == 0 && (visited[cr]&(1<<(cc+1))) == 0 {
						visited[cr] |= (1 << (cc + 1))
						stack[top] = uint16(cr<<8 | (cc + 1))
						top++
					}
				}

				unfillable += compSize % 4
				if unfillable > freeSpaces {
					return false
				}
			}
		}
	}
	return true
}
