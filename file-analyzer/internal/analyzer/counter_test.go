package analyzer

import "testing"

func TestCount(t *testing.T) {
	tests := []struct {
		name            string
		in              string
		wantWords       int
		wantLineBreaks  int
	}{
		{"empty", "", 0, 0},
		{"spaces only", "    \t   ", 0, 0},
		{"one word", "Hello", 1, 0},
		{"two words", "Hello world", 2, 0},
		{"newline end", "Hello\n", 1, 1},
		{"multi lines", "a b c\nd e\nf", 6, 2},
		{"tabs & mixed ws", "a\tb   c\n d", 4, 1},
		{"multiple newlines", "a\n\nb\n", 2, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotW, gotL := Count(tt.in)
			if gotW != tt.wantWords || gotL != tt.wantLineBreaks {
				t.Fatalf("Count(%q) = (%d,%d); want (%d,%d)",
					tt.in, gotW, gotL, tt.wantWords, tt.wantLineBreaks)
			}
		})
	}
}