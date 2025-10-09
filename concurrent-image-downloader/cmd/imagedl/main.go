package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/turman17/concurrent-image-downloader/internal/downloader"
)

func readLines(p string) ([]string, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var urls []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		urls = append(urls, line)
	}
	return urls, sc.Err()
}

func main() {
	input := flag.String("in", "./images.txt", "path to file with images")
	output := flag.String("out", "./downloads", "output directory")
	concurrency := flag.Int("c", 5, "number of concurrent downloads")
	timeout := flag.Duration("t", 10*time.Second, "per-request timeout (e.g. 10s, 1m)")
	flag.Parse()

	if *concurrency <= 0 || *timeout <= 0 {
		log.Fatalf("concurrency and timeout must be positive")
	}

	urls, err := readLines(*input)
	if err != nil {
		log.Fatalf("reading %s: %v", *input, err)
	}
	if len(urls) == 0 {
		log.Println("no URLs found — nothing to do")
		return
	}

	if !filepath.IsAbs(*output) {
		cwd, _ := os.Getwd()
		*output = filepath.Join(cwd, *output)
	}

	cfg := downloader.Config{
		Input:       *input,
		Output:      *output,
		Concurrency: *concurrency,
		Timeout:     *timeout,
	}

	dl, err := downloader.New(cfg)
	if err != nil {
		log.Fatalf("init downloader: %v", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	start := time.Now()
	results := dl.Run(ctx, urls)
	elapsed := time.Since(start).Truncate(time.Millisecond)

	var okCount, failCount int
	var totalBytes int64
	for _, r := range results {
		if r.Err != nil {
			failCount++
			continue
		}
		okCount++
		totalBytes += r.Bytes
	}

	fmt.Printf("\nDone in %s | success: %d, failed: %d, bytes: %d\n",
		elapsed, okCount, failCount, totalBytes)
}