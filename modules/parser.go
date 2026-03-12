package input

import (
	"strings"
)

// ParseTetrominoes splits input content into individual 4x4 tetromino blocks
func ParseTetrominoes(content string) ([]string, error) {
	if content == "" {
		return nil, &ParseError{"empty input"}
	}

	// Check for leading newlines
	if strings.HasPrefix(content, "\n") {
		return nil, &ParseError{"leading newlines not allowed"}
	}

	// Remove single trailing newline if present
	if strings.HasSuffix(content, "\n") {
		content = strings.TrimSuffix(content, "\n")
	}

	// Check for multiple trailing newlines after removing one
	if strings.HasSuffix(content, "\n") {
		return nil, &ParseError{"multiple trailing newlines"}
	}

	// Split by double newlines (separator between tetrominoes)
	blocks := strings.Split(content, "\n\n")

	if len(blocks) == 0 {
		return nil, &ParseError{"no tetrominoes found"}
	}

	var tetrominoes []string
	for _, block := range blocks {
		// Check for empty blocks (multiple separators)
		if block == "" {
			return nil, &ParseError{"multiple separators found"}
		}

		lines := strings.Split(block, "\n")

		// Each tetromino must have exactly 4 lines
		if len(lines) != 4 {
			return nil, &ParseError{"invalid tetromino format"}
		}

		// Validate each line has exactly 4 characters
		for _, line := range lines {
			if len(line) != 4 {
				return nil, &ParseError{"invalid line length"}
			}
		}

		tetrominoes = append(tetrominoes, block)
	}

	return tetrominoes, nil
}

// ParseError represents parsing errors
type ParseError struct {
	Message string
}

func (e *ParseError) Error() string {
	return e.Message
}
