package downloader

import "time"

type Config struct {
	Input       string
	Output      string
	Concurrency int
	Timeout     time.Duration
}

type Job struct {
	Idx int
	URL string
}

type Result struct {
	Idx       int
	URL       string
	Filename  string
	Bytes     int64
	Err       error
	Attempted int
}