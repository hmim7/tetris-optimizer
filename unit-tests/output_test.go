package unittests

import (
	"strings"
	"testing"
	modules "tetris-optimizer/modules"
)

func TestRenderOutput_BasicFormatting(t *testing.T) {
	lines := []string{
		"AB..",
		"A.B.",
		"....",
	}
	got := modules.RenderOutput(lines)
	expected := "AB..\nA.B.\n....\n"
	if got != expected {
		t.Fatalf("expected %q, got %q", expected, got)
	}
}

func TestRenderOutput_TrimsTrailingSpaces(t *testing.T) {
	lines := []string{
		"A..   ",
		"..B\t",
	}
	got := modules.RenderOutput(lines)
	expected := "A..\n..B\n"
	if got != expected {
		t.Fatalf("expected %q, got %q", expected, got)
	}

	// Ensure no line ends with whitespace
	for _, line := range strings.Split(strings.TrimSuffix(got, "\n"), "\n") {
		if strings.HasSuffix(line, " ") || strings.HasSuffix(line, "\t") {
			t.Fatalf("line has trailing whitespace: %q", line)
		}
	}
}

func TestRenderOutput_Empty(t *testing.T) {
	got := modules.RenderOutput(nil)
	if got != "" {
		t.Fatalf("expected empty string, got %q", got)
	}
}
