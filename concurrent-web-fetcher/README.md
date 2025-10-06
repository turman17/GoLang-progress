# Concurrent Web Fetcher

## Introduction

This is my third Go learning project, where I explore concurrency by building a simple web fetcher. The goal is to fetch multiple URLs concurrently, understand goroutines and channels, and get comfortable with handling HTTP requests in Go.

## What It Does

The concurrent web fetcher retrieves multiple web pages at the same time and prints their HTTP status codes along with the URLs. It demonstrates how to efficiently perform network I/O operations concurrently, improving overall performance.

## Project Structure

```
concurrent-web-fetcher/
├─ cmd/
│  └─ webfetch/
│     └─ main.go        # CLI entry point, argument parsing, printing results
├─ internal/
│  ├─ fetcher/
│  │  ├─ fetch.go       # Fetch(url, timeout) → Result
│  │  ├─ runner.go      # Run(urls, opts) with concurrency + timeout
│  │  └─ types.go       # Result struct (duration, status, bytes, error)
│  └─ input/
│     ├─ input.go       # LinkSplitter: reads URLs from file and filters invalid ones
│     └─ urls.txt       # Example file with 10 URLs (8 valid, 2 invalid)
└─ README.md
```

## Usage

Run the fetcher with:

```bash
go run main.go internal/input/urls.txt
```

### Sample Output

```
https://golang.org - 200 OK
https://example.com - 200 OK
https://invalid-url - error: Get "https://invalid-url": dial tcp: lookup invalid-url: no such host
... (other URLs)
```

## Learning Goals

- Understand and use goroutines for concurrent execution.
- Communicate between goroutines using channels.
- Handle HTTP requests and responses in Go.
- Implement timeout and error handling in concurrent programs.

## Next Steps

- Extend the fetcher to accept URLs from files or command-line arguments.
- Add timeout handling for HTTP requests to avoid hanging.
- Aggregate and summarize results after all fetches complete.
- Explore advanced concurrency patterns such as worker pools and context cancellation.

---

**Learning Journey:** This project is a stepping stone to mastering Go's concurrency model and building efficient networked applications.
