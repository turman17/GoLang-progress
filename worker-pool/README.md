# Worker Pool in Go

## Purpose

This project demonstrates the implementation of a Worker Pool pattern in Go. It is designed to efficiently manage and distribute tasks among a fixed number of worker goroutines, optimizing resource usage and improving concurrent processing.

## What It Does

The Worker Pool creates a set number of worker goroutines that listen on a job channel for incoming tasks. Each worker processes jobs independently and concurrently. The main routine dispatches tasks into the channel, and a sync.WaitGroup ensures all jobs are completed before the program exits.

## How to Run

1. Clone the repository:

```bash
git clone https://github.com/yourusername/worker-pool.git
cd worker-pool
```

2. Build and run the project:

```bash
go run main.go
```

## Example Command and Output

```bash
go run main.go
```

Example output:

```
Worker 1 started job 1
Worker 2 started job 2
Worker 3 started job 3
Worker 1 finished job 1
Worker 2 finished job 2
Worker 3 finished job 3
All jobs completed.
```

## Learning Goals

- Understand how to create and manage goroutines for concurrent execution.
- Use channels for communication and synchronization between goroutines.
- Employ sync.WaitGroup to wait for a collection of goroutines to finish.
- Implement the Worker Pool pattern to efficiently handle multiple tasks concurrently in Go.
