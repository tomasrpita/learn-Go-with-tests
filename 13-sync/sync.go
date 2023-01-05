package syncronicity

import "sync"

// Counter is a simple counter.
type Counter struct {
	mu    sync.Mutex // guards value
	value int
}

// Inc increments the counter.
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current value of the counter.
func (c *Counter) Value() int {
	return c.value
}

// NewCounter returns a new Counter.
func NewCounter() *Counter {
	return &Counter{}
}
