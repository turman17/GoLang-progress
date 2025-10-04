package main

import (
	"fmt"
	"os"

	"github.com/turman17/GoLang-progress/file-analyzer/internal/analyzer"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s <file>\n", os.Args[0])
	os.Exit(2)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}
	path := os.Args[1]

	words, lines, err := analyzer.ReadAndCount(path)
	if err != nil {
		// Print the error and exit non-zero
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Words: %d | Line breaks: %d\n", words, lines)
}