package unittests

import (
	"testing"
	"tetris-optimizer/modules"
)

func TestNewTetromino_Normalization(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []input.Coordinate
	}{
		{
			name:  "Square at origin",
			input: "##..\n##..\n....\n....",
			expected: []input.Coordinate{
				{Row: 0, Col: 0}, {Row: 0, Col: 1},
				{Row: 1, Col: 0}, {Row: 1, Col: 1},
			},
		},
		{
			name:  "Square offset - should normalize to origin",
			input: "....\n.##.\n.##.\n....",
			expected: []input.Coordinate{
				{Row: 0, Col: 0}, {Row: 0, Col: 1},
				{Row: 1, Col: 0}, {Row: 1, Col: 1},
			},
		},
		{
			name:  "I-shape horizontal",
			input: "....\n####\n....\n....",
			expected: []input.Coordinate{
				{Row: 0, Col: 0}, {Row: 0, Col: 1},
				{Row: 0, Col: 2}, {Row: 0, Col: 3},
			},
		},
		{
			name:  "I-shape vertical",
			input: "#...\n#...\n#...\n#...",
			expected: []input.Coordinate{
				{Row: 0, Col: 0}, {Row: 1, Col: 0},
				{Row: 2, Col: 0}, {Row: 3, Col: 0},
			},
		},
		{
			name:  "L-shape",
			input: "#...\n#...\n##..\n....",
			expected: []input.Coordinate{
				{Row: 0, Col: 0}, {Row: 1, Col: 0},
				{Row: 2, Col: 0}, {Row: 2, Col: 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tetromino := input.NewTetromino(tt.input)

			if len(tetromino.Blocks) != len(tt.expected) {
				t.Errorf("Expected %d blocks, got %d", len(tt.expected), len(tetromino.Blocks))
				return
			}

			// Check each coordinate
			for i, expected := range tt.expected {
				if tetromino.Blocks[i] != expected {
					t.Errorf("Block %d: expected %+v, got %+v", i, expected, tetromino.Blocks[i])
				}
			}
		})
	}
}

func TestTetromino_Dimensions(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedWidth  int
		expectedHeight int
	}{
		{
			name:           "Square 2x2",
			input:          "##..\n##..\n....\n....",
			expectedWidth:  2,
			expectedHeight: 2,
		},
		{
			name:           "I-shape horizontal 4x1",
			input:          "....\n####\n....\n....",
			expectedWidth:  4,
			expectedHeight: 1,
		},
		{
			name:           "I-shape vertical 1x4",
			input:          "#...\n#...\n#...\n#...",
			expectedWidth:  1,
			expectedHeight: 4,
		},
		{
			name:           "L-shape 2x3",
			input:          "#...\n#...\n##..\n....",
			expectedWidth:  2,
			expectedHeight: 3,
		},
		{
			name:           "T-shape 3x2",
			input:          "###.\n.#..\n....\n....",
			expectedWidth:  3,
			expectedHeight: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tetromino := input.NewTetromino(tt.input)

			if tetromino.Width() != tt.expectedWidth {
				t.Errorf("Expected width %d, got %d", tt.expectedWidth, tetromino.Width())
			}

			if tetromino.Height() != tt.expectedHeight {
				t.Errorf("Expected height %d, got %d", tt.expectedHeight, tetromino.Height())
			}
		})
	}
}

func TestTetromino_CanPlaceAt(t *testing.T) {
	// Create a simple 2x2 square tetromino
	tetromino := input.NewTetromino("##..\n##..\n....\n....")

	tests := []struct {
		name      string
		boardSize int
		startRow  int
		startCol  int
		expected  bool
	}{
		{
			name:      "Valid placement at origin",
			boardSize: 4,
			startRow:  0,
			startCol:  0,
			expected:  true,
		},
		{
			name:      "Valid placement at bottom-right corner",
			boardSize: 4,
			startRow:  2,
			startCol:  2,
			expected:  true,
		},
		{
			name:      "Invalid placement - exceeds right boundary",
			boardSize: 4,
			startRow:  0,
			startCol:  3,
			expected:  false,
		},
		{
			name:      "Invalid placement - exceeds bottom boundary",
			boardSize: 4,
			startRow:  3,
			startCol:  0,
			expected:  false,
		},
		{
			name:      "Invalid placement - negative position",
			boardSize: 4,
			startRow:  -1,
			startCol:  0,
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tetromino.CanPlaceAt(tt.boardSize, tt.startRow, tt.startCol)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestTetromino_GetAbsoluteCoordinates(t *testing.T) {
	// Create a simple L-shape tetromino
	tetromino := input.NewTetromino("#...\n#...\n##..\n....")

	tests := []struct {
		name     string
		startRow int
		startCol int
		expected []input.Coordinate
	}{
		{
			name:     "Placement at origin",
			startRow: 0,
			startCol: 0,
			expected: []input.Coordinate{
				{Row: 0, Col: 0}, {Row: 1, Col: 0},
				{Row: 2, Col: 0}, {Row: 2, Col: 1},
			},
		},
		{
			name:     "Placement offset",
			startRow: 2,
			startCol: 3,
			expected: []input.Coordinate{
				{Row: 2, Col: 3}, {Row: 3, Col: 3},
				{Row: 4, Col: 3}, {Row: 4, Col: 4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coords := tetromino.GetAbsoluteCoordinates(tt.startRow, tt.startCol)

			if len(coords) != len(tt.expected) {
				t.Errorf("Expected %d coordinates, got %d", len(tt.expected), len(coords))
				return
			}

			for i, expected := range tt.expected {
				if coords[i] != expected {
					t.Errorf("Coordinate %d: expected %+v, got %+v", i, expected, coords[i])
				}
			}
		})
	}
}

func TestTetromino_EmptyInput(t *testing.T) {
	tetromino := input.NewTetromino("....\n....\n....\n....")

	if len(tetromino.Blocks) != 0 {
		t.Errorf("Expected 0 blocks for empty input, got %d", len(tetromino.Blocks))
	}

	if tetromino.Width() != 0 {
		t.Errorf("Expected width 0 for empty tetromino, got %d", tetromino.Width())
	}

	if tetromino.Height() != 0 {
		t.Errorf("Expected height 0 for empty tetromino, got %d", tetromino.Height())
	}
}
