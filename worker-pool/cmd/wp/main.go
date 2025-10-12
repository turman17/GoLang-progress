package main

import (
	"flag"
	"log"
	"github.com/turman17/worker-pool/internal/pool"
)

func main(){
	workers := flag.Int("workers", 4, "number of workers")
	n := flag.Int("n", 20, "number of tasks")
	sleep := flag.Int("sleep", 10, "extra time for a task, in ms")
	flag.Parse()

	if *workers <= 0 || *n <= 0 || *sleep < 0 {
		log.Fatalf("Workers and n and sleep should be more than 0")
	}
	
	cfg := pool.Options{
		Workers: *workers,
		NTasks: *n,
		Sleep: *sleep,
	}

	pool.Run(cfg)
}