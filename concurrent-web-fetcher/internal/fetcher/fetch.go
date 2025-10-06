package fetcher

import (
	"io"
	"net/http"
	"time"
)

func Fetch(url string, timeout time.Duration) Result {
	start := time.Now()

	client := &http.Client{Timeout: timeout}
	resp, err := client.Get(url)
	if err != nil {
		return Result{
			URL:      url,
			Duration: time.Since(start),
			Err:      err,
		}
	}
	defer resp.Body.Close()

	n, _ := io.Copy(io.Discard, resp.Body)

	return Result{
		URL:      url,
		Duration: time.Since(start),
		Status:   resp.StatusCode,
		Bytes:    int(n),
		Err:      nil,
	}
}