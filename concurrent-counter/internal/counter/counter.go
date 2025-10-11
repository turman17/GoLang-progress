package counter

import "sync"

type Counter struct{
	mu sync.Mutex
	v int64
}

func (c *Counter) Inc(){
	c.mu.Lock()
	c.v++
	c.mu.Unlock()
}

func (c *Counter) Add(n int64){
	c.mu.Lock()
	c.v += n
	c.mu.Unlock()
}

func (c *Counter) Value() int64{
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v
}

type UnsafeCounter struct { v int64 }

func (c *UnsafeCounter) Inc(){c.v++}
func (c *UnsafeCounter) Add(n int64){c.v += n}
func (c *UnsafeCounter) Value() int64{return c.v}
