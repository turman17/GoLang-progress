package main

import (
	"flag"
	"fmt"
	"github.com/turman17/concurrent-counter/internal/counter"
)

func main() {
	workers := flag.Int("workers", 10, "Number of goroutines (concurrent workers) that will increment the counter.")
	iters := flag.Int("iters", 1000, "Number of increments each goroutine will perform.")
	start := flag.Int64("start", 0, "Initial value of the counter before the goroutines start.")
	unsafe := flag.Bool("unsafe", false, "If true, disables the mutex lock to demonstrate race conditions (unsafe mode).")
	flag.Parse()

	cfg := counter.Flags{
		Workers: *workers,
		Iters:   *iters,
		Start:   *start,
		Unsafe:  *unsafe,
	}

	res := counter.Run(cfg)
	
	fmt.Printf("workers=%d iters=%d start=%d unsafe=%v\n",
		cfg.Workers, cfg.Iters, cfg.Start, cfg.Unsafe)
	fmt.Printf("final=%d expected=%d elapsed=%s\n",
		res.Final, res.Expected(), res.Elapsed)
}
