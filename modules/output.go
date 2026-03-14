package modules

import "strings"

// RenderOutput converts a solved board (slice of strings) into the final output
// format with newline termination and no trailing whitespace.
func RenderOutput(lines []string) string {
	if len(lines) == 0 {
		return ""
	}

	var b strings.Builder
	for i, line := range lines {
		trimmed := strings.TrimRight(line, " \t")
		b.WriteString(trimmed)
		if i < len(lines)-1 {
			b.WriteByte('\n')
		}
	}
	b.WriteByte('\n')
	return b.String()
}
