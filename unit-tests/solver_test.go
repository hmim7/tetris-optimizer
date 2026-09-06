package unittests

import (
	"path/filepath"
	"testing"
	modules "tetris-optimizer/modules"
)

func TestSolve_GoldenSizesAndDots(t *testing.T) {
	tests := []struct {
		name         string
		filename     string
		expectedSize int
		expectedDots int
		expectedTets int
	}{
		{
			name:         "goodexample00",
			filename:     "goodexample00.txt",
			expectedSize: 2,
			expectedDots: 0,
			expectedTets: 1,
		},
		{
			name:         "goodexample01",
			filename:     "goodexample01.txt",
			expectedSize: 5,
			expectedDots: 9,
			expectedTets: 4,
		},
		{
			name:         "goodexample02",
			filename:     "goodexample02.txt",
			expectedSize: 6,
			expectedDots: 4,
			expectedTets: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content, err := modules.ReadFile(filepath.Join("..", "samples", tt.filename))
			if err != nil {
				t.Fatalf("ReadFile() error: %v", err)
			}

			blocks, err := modules.ParseTetrominoes(content)
			if err != nil {
				t.Fatalf("ParseTetrominoes() error: %v", err)
			}

			if len(blocks) != tt.expectedTets {
				t.Fatalf("expected %d tetrominoes, got %d", tt.expectedTets, len(blocks))
			}

			if err := modules.ValidateTetrominoes(blocks); err != nil {
				t.Fatalf("ValidateTetrominoes() error: %v", err)
			}

			tetrominoes := make([]*modules.Tetromino, 0, len(blocks))
			for _, b := range blocks {
				tetrominoes = append(tetrominoes, modules.NewTetromino(b))
			}

			solution, err := modules.Solve(tetrominoes)
			if err != nil {
				t.Fatalf("Solve() error: %v", err)
			}

			if len(solution) != tt.expectedSize {
				t.Fatalf("expected %dx%d output, got %dx?", tt.expectedSize, tt.expectedSize, len(solution))
			}

			dots := 0
			for _, line := range solution {
				if len(line) != tt.expectedSize {
					t.Fatalf("expected line length %d, got %d", tt.expectedSize, len(line))
				}
				for i := 0; i < len(line); i++ {
					if line[i] == '.' {
						dots++
					}
				}
			}

			if dots != tt.expectedDots {
				t.Fatalf("expected %d dots, got %d", tt.expectedDots, dots)
			}
		})
	}
}

func TestSolve_SampleTxt_ExactOutput(t *testing.T) {
	content, err := modules.ReadFile(filepath.Join("..", "samples", "sample.txt"))
	if err != nil {
		t.Fatalf("ReadFile() error: %v", err)
	}

	blocks, err := modules.ParseTetrominoes(content)
	if err != nil {
		t.Fatalf("ParseTetrominoes() error: %v", err)
	}

	if err := modules.ValidateTetrominoes(blocks); err != nil {
		t.Fatalf("ValidateTetrominoes() error: %v", err)
	}

	tetrominoes := make([]*modules.Tetromino, 0, len(blocks))
	for _, b := range blocks {
		tetrominoes = append(tetrominoes, modules.NewTetromino(b))
	}

	solution, err := modules.Solve(tetrominoes)
	if err != nil {
		t.Fatalf("Solve() error: %v", err)
	}

	expected := []string{
		"ABBBB.",
		"ACCCEE",
		"AFFCEE",
		"A.FFGG",
		"HHHDDG",
		".HDD.G",
	}

	if len(solution) != len(expected) {
		t.Fatalf("expected %d lines, got %d", len(expected), len(solution))
	}

	for i := range expected {
		if solution[i] != expected[i] {
			t.Fatalf("line %d: expected %q, got %q", i, expected[i], solution[i])
		}
	}
}
