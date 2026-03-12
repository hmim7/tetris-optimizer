package unittests

import (
	"testing"
	"tetris-optimizer/modules"
)

func TestParseTetrominoes_ValidInput(t *testing.T) {
	testInput := "....\n.##.\n.##.\n...."

	result, err := input.ParseTetrominoes(testInput)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("Expected 1 tetromino, got: %d", len(result))
	}

	if result[0] != testInput {
		t.Errorf("Expected %q, got: %q", testInput, result[0])
	}
}

func TestParseTetrominoes_MultipleBlocks(t *testing.T) {
	testInput := "...#\n...#\n...#\n...#\n\n....\n....\n....\n####"

	result, err := input.ParseTetrominoes(testInput)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("Expected 2 tetrominoes, got: %d", len(result))
	}

	expected1 := "...#\n...#\n...#\n...#"
	expected2 := "....\n....\n....\n####"

	if result[0] != expected1 {
		t.Errorf("Expected first tetromino %q, got: %q", expected1, result[0])
	}

	if result[1] != expected2 {
		t.Errorf("Expected second tetromino %q, got: %q", expected2, result[1])
	}
}

func TestParseTetrominoes_EmptyInput(t *testing.T) {
	_, err := input.ParseTetrominoes("")

	if err == nil {
		t.Error("Expected error for empty input")
	}
}

func TestParseTetrominoes_LeadingNewline(t *testing.T) {
	testInput := "\n....\n.##.\n.##.\n...."

	_, err := input.ParseTetrominoes(testInput)

	if err == nil {
		t.Error("Expected error for leading newline")
	}
}

func TestParseTetrominoes_MultipleSeparators(t *testing.T) {
	testInput := "...#\n...#\n...#\n...#\n\n\n....\n....\n....\n####"

	_, err := input.ParseTetrominoes(testInput)

	if err == nil {
		t.Error("Expected error for multiple separators")
	}
}

func TestParseTetrominoes_InvalidLineCount(t *testing.T) {
	testInput := "...#\n...#\n...#"

	_, err := input.ParseTetrominoes(testInput)

	if err == nil {
		t.Error("Expected error for invalid line count")
	}
}

func TestParseTetrominoes_InvalidLineLength(t *testing.T) {
	testInput := "...\n...#\n...#\n...#"

	_, err := input.ParseTetrominoes(testInput)

	if err == nil {
		t.Error("Expected error for invalid line length")
	}
}

func TestParseTetrominoes_MissingSeparator(t *testing.T) {
	testInput := "...#\n...#\n...#\n...#\n....\n....\n....\n####"

	_, err := input.ParseTetrominoes(testInput)

	if err == nil {
		t.Error("Expected error for missing separator")
	}
}

func TestParseTetrominoes_SingleTrailingNewline(t *testing.T) {
	testInput := "....\n.##.\n.##.\n....\n"

	result, err := input.ParseTetrominoes(testInput)

	if err != nil {
		t.Errorf("Expected no error for single trailing newline, got: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("Expected 1 tetromino, got: %d", len(result))
	}
}

func TestParseTetrominoes_MultipleTrailingNewlines(t *testing.T) {
	testInput := "....\n.##.\n.##.\n....\n\n"

	_, err := input.ParseTetrominoes(testInput)

	if err == nil {
		t.Error("Expected error for multiple trailing newlines")
	}
}

func TestParseTetrominoes_HalfBlockEOF(t *testing.T) {
	testInput := "...#\n...#\n...#\n...#\n\n....\n...."

	_, err := input.ParseTetrominoes(testInput)

	if err == nil {
		t.Error("Expected error for incomplete block at EOF")
	}
}
