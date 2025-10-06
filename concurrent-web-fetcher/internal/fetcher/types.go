package fetcher

import "time"

type Result struct {
    URL      string
    Duration time.Duration
    Status   int
    Bytes    int
    Err      error
}
