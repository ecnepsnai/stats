package stats

import "sync"

// Counter describes a simple, thread-safe counter. Counters have an initial value of 0 and a maximum value of
// UINT64_MAX. Counter values cannot be negative.
type Counter struct {
	value uint64
	lock  *sync.RWMutex
}

// NewCounter creates a new counter with a default value of 0.
func NewCounter() *Counter {
	return &Counter{
		value: 0,
		lock:  &sync.RWMutex{},
	}
}

// Increment will increment the value of the counter by 1.
func (c *Counter) Increment() {
	c.IncrementBy(1)
}

// Decrement will decrement the value of the counter by 1. If the counter's current value is 0, does nothing.
func (c *Counter) Decrement() {
	c.DecrementBy(1)
}

// IncrementBy will increment the value of the counter by n.
func (c *Counter) IncrementBy(n uint64) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.value += n
}

// DecrementBy will decrement the value of the counter by n. If the counter's current value is less than n,
// sets the value to 0.
func (c *Counter) DecrementBy(n uint64) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.value < n {
		c.value = 0
	} else {
		c.value -= n
	}
}

// Set will update the value of the counter to n
func (c *Counter) Set(n uint64) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.value = n
}

// Get will return the current value of the counter
func (c *Counter) Get() uint64 {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.value
}

// GetAndSet will update the value of this counter to n and return the previous value from before the update
func (c *Counter) GetAndSet(n uint64) uint64 {
	c.lock.Lock()
	defer c.lock.Unlock()

	v := c.value
	c.value = n
	return v
}
