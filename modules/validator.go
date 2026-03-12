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

	// Validate each line
	for _, line := range lines {
		// Validate line length is exactly 4
		if len(line) != 4 {
			return &ValidationError{"invalid line length"}
		}

		// Validate characters and count blocks
		for _, char := range line {
			if char != '.' && char != '#' {
				return &ValidationError{"invalid character"}
			}
			if char == '#' {
				blockCount++
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

	return nil
}

// ValidationError represents validation errors
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}
