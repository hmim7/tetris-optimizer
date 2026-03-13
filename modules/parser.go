package modules

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

	// Remove single trailing newline if present (TrimSuffix handles the check automatically)
	content = strings.TrimSuffix(content, "\n")

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

		if err := validateBlockDimensions(block); err != nil {
			return nil, err
		}

		tetrominoes = append(tetrominoes, block)
	}

	return tetrominoes, nil
}

func validateBlockDimensions(block string) error {
	lines := strings.Split(block, "\n")
	if len(lines) != 4 {
		return &ParseError{"invalid line count"}
	}
	for _, line := range lines {
		if len(line) != 4 {
			return &ParseError{"invalid line length"}
		}
	}
	return nil
}

// ParseError represents parsing errors
type ParseError struct {
	Message string
}

func (e *ParseError) Error() string {
	return e.Message
}
