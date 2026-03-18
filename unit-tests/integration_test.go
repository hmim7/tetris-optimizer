package unittests

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	modules "tetris-optimizer/modules"
	"time"
)

func runPipeline(args []string) (string, error) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	os.Args = args
	content, err := modules.ProcessInput()
	if err != nil {
		return "ERROR", err
	}

	blocks, err := modules.ParseTetrominoes(content)
	if err != nil {
		return "ERROR", err
	}

	if err := modules.ValidateTetrominoes(blocks); err != nil {
		return "ERROR", err
	}

	tetrominoes := make([]*modules.Tetromino, 0, len(blocks))
	for _, b := range blocks {
		tetrominoes = append(tetrominoes, modules.NewTetromino(b))
	}

	solution, err := modules.Solve(tetrominoes)
	if err != nil {
		return "ERROR", err
	}

	return modules.RenderOutput(solution), nil
}

func splitOutput(output string) []string {
	trimmed := strings.TrimSuffix(output, "\n")
	if trimmed == "" {
		return nil
	}
	return strings.Split(trimmed, "\n")
}

func countDots(lines []string) int {
	dots := 0
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			if line[i] == '.' {
				dots++
			}
		}
	}
	return dots
}

func verifyLetterAssignment(t *testing.T, lines []string, tetroCount int) {
	t.Helper()

	counts := make(map[byte]int)
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			ch := line[i]
			if ch >= 'A' && ch <= 'Z' {
				counts[ch]++
			}
		}
	}

	if len(counts) != tetroCount {
		t.Fatalf("expected %d letters, got %d", tetroCount, len(counts))
	}

	for i := 0; i < tetroCount; i++ {
		letter := byte('A' + i)
		if counts[letter] != 4 {
			t.Fatalf("expected letter %c to appear 4 times, got %d", letter, counts[letter])
		}
	}
}

func TestIntegration_CoreFunctionality(t *testing.T) {
	tests := []struct {
		name         string
		file         string
		size         int
		dots         int
		tetrominoes  int
		expectOutput bool
	}{
		{"goodexample00", "goodexample00.txt", 2, 0, 1, true},
		{"goodexample01", "goodexample01.txt", 5, 9, 4, true},
		{"goodexample02", "goodexample02.txt", 6, 4, 8, true},
		{"goodexample03", "goodexample03.txt", 7, 5, 11, true},
		{"hardexam", "hardexam.txt", 7, 1, 12, true},
		{"non_standard_placement", "non_standard_placement.txt", 3, 5, 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := runPipeline([]string{"program", filepath.Join("..", "samples", tt.file)})
			if err != nil {
				t.Fatalf("runPipeline error: %v", err)
			}

			lines := splitOutput(output)
			if len(lines) != tt.size {
				t.Fatalf("expected %dx%d output, got %dx?", tt.size, tt.size, len(lines))
			}
			for _, line := range lines {
				if len(line) != tt.size {
					t.Fatalf("expected line length %d, got %d", tt.size, len(line))
				}
				if strings.HasSuffix(line, " ") || strings.HasSuffix(line, "\t") {
					t.Fatalf("line has trailing whitespace: %q", line)
				}
			}

			if dots := countDots(lines); dots != tt.dots {
				t.Fatalf("expected %d dots, got %d", tt.dots, dots)
			}

			verifyLetterAssignment(t, lines, tt.tetrominoes)
		})
	}
}

func TestIntegration_DeterministicOutput(t *testing.T) {
	path := filepath.Join("..", "samples", "goodexample01.txt")
	output1, err := runPipeline([]string{"program", path})
	if err != nil {
		t.Fatalf("runPipeline error: %v", err)
	}
	output2, err := runPipeline([]string{"program", path})
	if err != nil {
		t.Fatalf("runPipeline error: %v", err)
	}
	if output1 != output2 {
		t.Fatalf("expected deterministic output, got different results")
	}
}

func TestIntegration_ErrorCases(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"missing_args", []string{"program"}},
		{"too_many_args", []string{"program", "file1.txt", "file2.txt"}},
		{"empty_file", []string{"program", filepath.Join("..", "samples", "empty_file.txt")}},
		{"invalid_chars_tabs", []string{"program", filepath.Join("..", "samples", "invalid_chars_tabs.txt")}},
		{"invalid_chars_spaces", []string{"program", filepath.Join("..", "samples", "invalid_chars_spaces.txt")}},
		{"invalid_chars_numbers", []string{"program", filepath.Join("..", "samples", "invalid_chars_numbers.txt")}},
		{"line_too_long", []string{"program", filepath.Join("..", "samples", "incorrect_line_length_long.txt")}},
		{"line_too_short", []string{"program", filepath.Join("..", "samples", "incorrect_line_length_short.txt")}},
		{"block_too_tall", []string{"program", filepath.Join("..", "samples", "incorrect_block_height_long.txt")}},
		{"block_too_short", []string{"program", filepath.Join("..", "samples", "incorrect_block_height_short.txt")}},
		{"missing_separator", []string{"program", filepath.Join("..", "samples", "missing_separator.txt")}},
		{"multiple_separators", []string{"program", filepath.Join("..", "samples", "badformat.txt")}},
		{"leading_newline", []string{"program", filepath.Join("..", "samples", "leading_newline.txt")}},
		{"trailing_newlines", []string{"program", filepath.Join("..", "samples", "trailing_newlines.txt")}},
		{"non_uniform_blocks", []string{"program", filepath.Join("..", "samples", "non_uniform_blocks.txt")}},
		{"half_block_eof", []string{"program", filepath.Join("..", "samples", "half_block_eof.txt")}},
		{"ghost_block_trailing", []string{"program", filepath.Join("..", "samples", "ghost_block_trailing.txt")}},
		{"too_many_blocks", []string{"program", filepath.Join("..", "samples", "badexample00.txt")}},
		{"too_few_blocks", []string{"program", filepath.Join("..", "samples", "too_few_blocks.txt")}},
		{"zero_blocks", []string{"program", filepath.Join("..", "samples", "badexample03.txt")}},
		{"diagonal_only", []string{"program", filepath.Join("..", "samples", "badexample01.txt")}},
		{"disconnected_pairs", []string{"program", filepath.Join("..", "samples", "badexample04.txt")}},
		{"vertical_gap", []string{"program", filepath.Join("..", "samples", "badexample02.txt")}},
		{"only_dots", []string{"program", filepath.Join("..", "samples", "only_dots.txt")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := runPipeline(tt.args)
			if err == nil {
				t.Fatalf("expected error, got nil with output %q", output)
			}
			if output != "ERROR" {
				t.Fatalf("expected output ERROR, got %q", output)
			}
		})
	}
}

func TestIntegration_PerformanceLimits(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		limit    time.Duration
		tetromos int
	}{
		{"8_tetros", "goodexample02.txt", 1 * time.Second, 8},
		{"11_tetros", "goodexample03.txt", 3 * time.Second, 11},
		{"12_tetros", "hardexam.txt", 5 * time.Second, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			_, err := runPipeline([]string{"program", filepath.Join("..", "samples", tt.file)})
			if err != nil {
				t.Fatalf("runPipeline error: %v", err)
			}
			elapsed := time.Since(start)
			if elapsed > tt.limit {
				t.Fatalf("expected %d tetrominoes to finish within %s, took %s", tt.tetromos, tt.limit, elapsed)
			}
		})
	}
}
