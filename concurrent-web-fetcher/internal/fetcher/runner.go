package fetcher

import (
	"time"
	"sync"
)

func Runner(links []string, timeout time.Duration, concurrency int) []Result {
	if concurrency <= 0{
		concurrency = 10
	}
	if timeout <= 0{
		timeout = 5 * time.Second
	}

	results := make([]Result, len(links))
	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrency)

	for i, u := range links {
		wg.Add(1)
		i, u := i, u
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			results[i] = Fetch(u, timeout)
		}()
	}

	wg.Wait()
	return results
}