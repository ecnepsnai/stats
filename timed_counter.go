package stats

import (
	"sync"
	"time"
)

// TimedCounter describes a timed counter object. A TimedCounter operates similarly to the Counter, except that it
// records each time the counter was incremented. This allows for easy tracking of statistics across periods of time,
// for example; "Number of failed requests in the last 24 hours".
type TimedCounter struct {
	points        []int64
	maxAge        time.Duration
	lock          *sync.RWMutex
	nextCleanupAt int64
}

// NewTimedCounter create a new TimedCounter with the defined maximum age of data points. If MaxAge is 0, there is no
// age defined and points are retained forever.
func NewTimedCounter(MaxAge time.Duration) *TimedCounter {
	return &TimedCounter{
		points: []int64{},
		maxAge: MaxAge,
		lock:   &sync.RWMutex{},
	}
}

// Increment will increment this timed counter
func (c *TimedCounter) Increment() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.points = append(c.points, time.Now().UTC().Unix())
	if len(c.points) == 1 {
		c.nextCleanupAt = time.Now().Unix() + int64(c.maxAge.Seconds())
	}
}

// DecrementLast will remove the last entry from this timed counter
func (c *TimedCounter) DecrementLast() {
	c.lock.Lock()
	defer c.lock.Unlock()

	count := len(c.points)
	if count == 0 {
		return
	} else if count == 1 {
		c.points = []int64{}
		c.nextCleanupAt = 0
	} else {
		c.points = c.points[0 : len(c.points)-1]
	}
}

// GetAll will return the current value of the counter
func (c *TimedCounter) GetAll() uint64 {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return uint64(len(c.points))
}

// GetLast will return the number of increments to the counter since time.Now()-d
func (c *TimedCounter) GetLast(d time.Duration) uint64 {
	c.lock.RLock()
	defer c.lock.RUnlock()

	t := uint64(0)
	for _, point := range c.points {
		if time.Since(time.Unix(point, 0).UTC()) <= d {
			t++
		}
	}

	return t
}

// Cleanup will remove expired points from this timed counter if needed. You should run Cleanup about as frequently as
// you expect to be incrementing the counter.
func (c *TimedCounter) Cleanup() {
	c.lock.Lock()
	defer c.lock.Unlock()

	if len(c.points) == 0 {
		return
	}
	if time.Now().UTC().Unix() < c.nextCleanupAt {
		return
	}

	i := len(c.points) - 1
	for i >= 0 {
		point := c.points[i]

		if time.Since(time.Unix(point, 0).UTC()) > c.maxAge {
			c.points = append(c.points[:i], c.points[i+1:]...)
		}

		i--
	}
}
