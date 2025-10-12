package pool

import (
	"fmt"
	"sync"
	"time"
)

func Run(cfg Options) {
	jobCh := make(chan Job)
	resCh := make(chan Result)

	var wg sync.WaitGroup
	wg.Add(cfg.Workers)
	for w := 0; w < cfg.Workers; w++ {
		go func() {
			defer wg.Done()
			for j := range jobCh {
				start := time.Now()

				time.Sleep(time.Duration(cfg.Sleep) * time.Millisecond)
				out := j.Value * j.Value

				resCh <- Result{
					JobID:  j.ID,
					Input:  j.Value,
					Output: out,
					time:   time.Since(start),
					Err:    nil,
				}
			}
		}()
	}
	go func() {
		for i := 0; i < cfg.NTasks; i++ {
			jobCh <- Job{ID: i, Value: i}
		}
		close(jobCh)
	}()

	go func() {
		wg.Wait()
		close(resCh)
	}()
	for r := range resCh {
		if r.Err != nil {
			fmt.Printf("job %d error: %v\n", r.JobID, r.Err)
			continue
		}
		fmt.Printf("job=%d in=%d out=%d took=%v\n", r.JobID, r.Input, r.Output, r.time)
	}
}
