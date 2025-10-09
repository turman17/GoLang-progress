# Concurrent Image Downloader

A simple Go project to practice concurrency by downloading images from a list of URLs in parallel using goroutines, channels, and the `os`/`io` packages.

## Features

- Reads a list of image URLs from a text file.
- Downloads images concurrently with configurable concurrency level.
- Saves images locally with unique filenames.
- Supports per-request timeout and retry logic.
- Prints summary of results: number of successes, failures, and total bytes downloaded.

## Project Structure

```
.
├── README.md
├── cmd
│   └── imagedl
│       └── main.go          # CLI entrypoint
├── go.mod
├── images.txt               # Example input with URLs
└── internal
    └── downloader
        ├── client.go        # HTTP client helper
        ├── downloader.go    # Core download logic
        ├── types.go         # Config, Job, Result types
        └── util.go          # Utility functions
```

## Usage

1. Build or run the CLI:
   ```bash
   go run ./cmd/imagedl -in=images.txt -out=downloads -c=5 -t=15s -r=2
   ```

2. Flags:
   - `-in` : Path to the input file containing image URLs (default `./images.txt`).
   - `-out`: Directory to save downloads (default `./downloads`).
   - `-c`  : Number of concurrent downloads (default `5`).
   - `-t`  : Per-request timeout (default `10s`).
   - `-r`  : Number of retries per URL (default `1`).

## Example

Given an `images.txt` file like:

```
https://example.com/image1.jpg
https://example.com/image2.png
```

Running:

```bash
go run ./cmd/imagedl -in=images.txt -out=downloads -c=8 -t=20s -r=3
```

Produces output like:

```
Done in 1.2s | success: 2, failed: 0, bytes: 4.3 MB
```

And downloaded files will be saved in the `downloads/` directory.

## Learning Goals

This project is meant to practice:

- Goroutines for concurrency.
- Buffered channels for job/result passing.
- Context cancellation and timeouts.
- Error handling and retries.
- Basic CLI with Go `flag` package.
