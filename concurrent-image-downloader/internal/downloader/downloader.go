package downloader

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"sync"
	"time"
)

type Downloader struct {
	cfg    Config
	client *http.Client
}

func New(cfg Config) (*Downloader, error) {
	if cfg.Concurrency <= 0 {
		cfg.Concurrency = 4
	}
	if cfg.Timeout <= 0 {
		cfg.Timeout = 30 * time.Second
	}
	if err := ensureDir(cfg.Output); err != nil {
		return nil, fmt.Errorf("create output dir: %w", err)
	}
	return &Downloader{
		cfg:    cfg,
		client: NewHTTPClient(cfg.Timeout),
	}, nil
}

func (d *Downloader) Run(ctx context.Context, urls []string) []Result {
	jobs := make(chan Job)
	results := make(chan Result, len(urls))

	var wg sync.WaitGroup
	for i := 0; i < d.cfg.Concurrency; i++ { // fixed: < not <=
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				res := d.downloadOnce(ctx, job)
				select {
				case results <- res:
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	// feed jobs
	go func() {
		for i, raw := range urls {
			norm, err := normalizeURL(raw)
			if err != nil {
				results <- Result{Idx: i, URL: raw, Err: err}
				continue
			}
			select {
			case jobs <- Job{Idx: i, URL: norm}:
			case <-ctx.Done():
				close(jobs)
				return
			}
		}
		close(jobs)
	}()

	// close results after workers complete
	go func() {
		wg.Wait()
		close(results)
	}()

	out := make([]Result, 0, len(urls))
	for r := range results {
		out = append(out, r)
	}
	return out
}

func (d *Downloader) downloadOnce(ctx context.Context, job Job) Result {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, job.URL, nil)
	if err != nil {
		return Result{Idx: job.Idx, URL: job.URL, Err: err}
	}
	resp, err := d.client.Do(req)
	if err != nil {
		return Result{Idx: job.Idx, URL: job.URL, Err: err}
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return Result{Idx: job.Idx, URL: job.URL, Err: fmt.Errorf("http %d", resp.StatusCode)}
	}

	name := fileNameFromURL(job.URL)
	full := filepath.Join(d.cfg.Output, name)
	n, err := writeFile(full, resp.Body)
	return Result{Idx: job.Idx, URL: job.URL, Filename: full, Bytes: n, Err: err}
}