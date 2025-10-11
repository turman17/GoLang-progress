package counter

import (
	"sync"
	"time"
)

func Run(cfg Flags) Result{
	startTime := time.Now()

	type ciface interface{
		Inc()
		Add(n int64)
		Value()int64
	}

	var c ciface
	if cfg.Unsafe{
		uc := &UnsafeCounter{}
		uc.Add(int64(cfg.Start))
		c = uc
	} else {
		sc := &Counter{}
		sc.Add(int64(cfg.Start))
		c = sc
	}
	var wg sync.WaitGroup

	for w := 0 ; w < cfg.Workers; w++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			for i := 0; i < cfg.Iters; i++{
				c.Inc()
			}
		}()
	}
	wg.Wait()

	return Result{
		Final:   c.Value(),
		Start:   cfg.Start,
		Workers: cfg.Workers,
		Iters:   cfg.Iters,
		Elapsed: time.Since(startTime),
	}
}