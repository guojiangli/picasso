package kcycle

import (
	"sync"
	"sync/atomic"
)

// Cycle ..
type Cycle struct {
	mu      *sync.Mutex
	wg      *sync.WaitGroup
	quit    chan error
	closing uint32
}

// NewCycle new a cycle life
func NewCycle() *Cycle {
	return &Cycle{
		mu:      &sync.Mutex{},
		wg:      &sync.WaitGroup{},
		quit:    make(chan error),
		closing: 0,
	}
}

// Run a new goroutine
func (c *Cycle) Run(fn func() error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.wg.Add(1)
	go func(c *Cycle) {
		defer c.wg.Done()
		if err := fn(); err != nil {
			if v := atomic.LoadUint32(&c.closing); v == 0 {
				c.quit <- err
			}
		}
	}(c)
}

// Close ..
func (c *Cycle) Close() {
	if atomic.CompareAndSwapUint32(&c.closing, 0, 1) {
		close(c.quit)
	}
}

// Wait blocked for a life cycle
func (c *Cycle) Wait() <-chan error {
	return c.quit
}
