package counter

import "time"

type Flags struct {
	Workers int
	Iters   int
	Start   int64
	Unsafe  bool
}

type Result struct {
	Final   int64
	Start   int64
	Workers int
	Iters   int
	Elapsed time.Duration
}

func (r Result) Expected() int64 {
	return r.Start + int64(r.Iters*r.Workers)
}
