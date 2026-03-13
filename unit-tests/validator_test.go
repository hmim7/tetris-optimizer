package unittests

import (
	"testing"
	"tetris-optimizer/modules"
)

func TestValidateTetromino_ValidBlock(t *testing.T) {
	validTetromino := "....\n.##.\n.##.\n...."

	err := modules.ValidateTetromino(validTetromino)

	if err != nil {
		t.Errorf("Expected no error for valid tetromino, got: %v", err)
	}
}

func TestValidateTetromino_InvalidCharacters(t *testing.T) {
	tests := []struct {
		name      string
		tetromino string
	}{
		{
			name:      "Numbers",
			tetromino: "...1\n...#\n...#\n...#",
		},
		{
			name:      "Spaces",
			tetromino: "... \n...#\n...#\n...#",
		},
		{
			name:      "Tabs",
			tetromino: "...\t\n...#\n...#\n...#",
		},
		{
			name:      "Letters",
			tetromino: "...A\n...#\n...#\n...#",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := modules.ValidateTetromino(tt.tetromino)
			if err == nil {
				t.Errorf("Expected error for invalid character in %s", tt.name)
			}
		})
	}
}

func TestValidateTetromino_InvalidLineLength(t *testing.T) {
	tests := []struct {
		name      string
		tetromino string
	}{
		{
			name:      "Line too short",
			tetromino: "...\n...#\n...#\n...#",
		},
		{
			name:      "Line too long",
			tetromino: ".....\n...#.\n...#.\n...#.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := modules.ValidateTetromino(tt.tetromino)
			if err == nil {
				t.Errorf("Expected error for invalid line length in %s", tt.name)
			}
		})
	}
}

func TestValidateTetromino_InvalidBlockHeight(t *testing.T) {
	tests := []struct {
		name      string
		tetromino string
	}{
		{
			name:      "Too few lines",
			tetromino: "...#\n...#\n...#",
		},
		{
			name:      "Too many lines",
			tetromino: "...#\n...#\n...#\n...#\n....",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := modules.ValidateTetromino(tt.tetromino)
			if err == nil {
				t.Errorf("Expected error for invalid block height in %s", tt.name)
			}
		})
	}
}

func TestValidateTetromino_InvalidBlockCount(t *testing.T) {
	tests := []struct {
		name      string
		tetromino string
	}{
		{
			name:      "Too few blocks",
			tetromino: "....\n.##.\n....\n....",
		},
		{
			name:      "Too many blocks",
			tetromino: "####\n####\n....\n....",
		},
		{
			name:      "No blocks",
			tetromino: "....\n....\n....\n....",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := modules.ValidateTetromino(tt.tetromino)
			if err == nil {
				t.Errorf("Expected error for invalid block count in %s", tt.name)
			}
		})
	}
}

func TestValidateTetrominoes_ValidMultiple(t *testing.T) {
	tetrominoes := []string{
		"....\n.##.\n.##.\n....",
		"...#\n...#\n...#\n...#",
		"....\n....\n....\n####",
	}

	err := modules.ValidateTetrominoes(tetrominoes)

	if err != nil {
		t.Errorf("Expected no error for valid tetrominoes, got: %v", err)
	}
}

func TestValidateTetrominoes_EmptyArray(t *testing.T) {
	var tetrominoes []string

	err := modules.ValidateTetrominoes(tetrominoes)

	if err == nil {
		t.Error("Expected error for empty tetrominoes array")
	}
}

func TestValidateTetrominoes_InvalidInArray(t *testing.T) {
	tetrominoes := []string{
		"....\n.##.\n.##.\n....",
		"...1\n...#\n...#\n...#", // Invalid character
		"....\n....\n....\n####",
	}

	err := modules.ValidateTetrominoes(tetrominoes)

	if err == nil {
		t.Error("Expected error for invalid tetromino in array")
	}
}

func TestValidateTetromino_ConnectedBlocks(t *testing.T) {
	tests := []struct {
		name      string
		tetromino string
		expected  bool
	}{
		{
			name:      "Valid L-shape",
			tetromino: "#...\n#...\n##..\n....",
			expected:  true,
		},
		{
			name:      "Valid I-shape",
			tetromino: "....\n####\n....\n....",
			expected:  true,
		},
		{
			name:      "Valid square",
			tetromino: "....\n.##.\n.##.\n....",
			expected:  true,
		},
		{
			name:      "Disconnected pairs",
			tetromino: "##..\n....\n..##\n....",
			expected:  false,
		},
		{
			name:      "Diagonal only",
			tetromino: "#...\n.#..\n..#.\n...#",
			expected:  false,
		},
		{
			name:      "Vertical gap",
			tetromino: "#...\n....\n#...\n##..",
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := modules.ValidateTetromino(tt.tetromino)
			if tt.expected && err != nil {
				t.Errorf("Expected valid tetromino for %s, got error: %v", tt.name, err)
			}
			if !tt.expected && err == nil {
				t.Errorf("Expected invalid tetromino for %s, got no error", tt.name)
			}
		})
	}
}
