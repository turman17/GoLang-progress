package main

import (
	"fmt"
	"os"
	"time"

	"github.com/turman17/concurrent-web-fetcher/internal/fetcher"
	"github.com/turman17/concurrent-web-fetcher/internal/input"
)

func error(msg string, exitCode int) {
	fmt.Println(msg, "\n")
	os.Exit(exitCode)
}

func main() {
	args := os.Args
	if len(args) != 2 {
		error("Error: Wrong number of urguments. ./web-fetcher <file-path/links>", 1)
	}

	links, err := input.LinkSplitter(args[1])
	if err != nil {
		error("Error: Splitting Links", 1)
	}
	res := fetcher.Runner(links, 10*time.Second, 10)

	ok, fail := 0, 0
	for _, r := range res {
		if r.Err != nil {
			fmt.Printf("%7s  ERROR: %v  %s\n", r.Duration, r.Err, r.URL)
			fail++
			continue
		}
		fmt.Printf("%7s  %d  %d B  %s\n", r.Duration, r.Status, r.Bytes, r.URL)
		ok++
	}
	fmt.Printf("Done (%d ok, %d error)\n", ok, fail)
}
