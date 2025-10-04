package analyzer

import "strings"

func Count(data string) (int, int) {
	words := len(strings.Fields(data))
	lines := strings.Count(data, "\n")
	return words, lines
}