package stats

import (
	"sync"
	"time"

	"github.com/ecnepsnai/ring"
)

// Timer describes an object which collects samples of timed operations.
type Timer struct {
	ring ring.Ring
	lock *sync.RWMutex
}

// NewTimer create a new timer that accepts a maximum number of samples.
func NewTimer(MaxSamples int) *Timer {
	return &Timer{
		ring: *ring.New(MaxSamples),
		lock: &sync.RWMutex{},
	}
}

// AddSample add a new sample to the timer. If the timer already has as many samples as defined as the maximum, the
// oldest sample is discarded.
func (t *Timer) AddSample(d time.Duration) {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.ring.Add(d)
}

// Reset will remove all samples from the timer.
func (t *Timer) Reset() {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.ring.Truncate()
}

// GetLast return the last sample added to the timer, or 0.
func (t *Timer) GetLast() time.Duration {
	t.lock.RLock()
	defer t.lock.RUnlock()

	last := t.ring.Last()
	if last == nil {
		return 0
	}

	d, ok := last.(time.Duration)
	if !ok {
		panic("object found in ring was not type time.Duration")
	}
	return d
}

// GetAll return all samples in the timer, or an empty slice.
func (t *Timer) GetAll() []time.Duration {
	t.lock.RLock()
	defer t.lock.RUnlock()

	objects := t.ring.All()
	if len(objects) == 0 {
		return []time.Duration{}
	}

	samples := make([]time.Duration, len(objects))
	for i, obj := range objects {
		d, ok := obj.(time.Duration)
		if !ok {
			panic("object found in ring was not type time.Duration")
		}
		samples[i] = d
	}

	return samples
}

// GetAverage return the average value from all samples in the timer.
func (t *Timer) GetAverage() time.Duration {
	samples := t.GetAll()
	if len(samples) == 0 {
		return 0
	}

	sum := int64(0)
	for _, sample := range samples {
		sum += int64(sample)
	}

	return time.Duration(sum / int64(len(samples)))
}

// GetMinimum return the minimum value from all samples in the timer.
func (t *Timer) GetMinimum() time.Duration {
	samples := t.GetAll()
	if len(samples) == 0 {
		return 0
	}

	min := ^uint64(0)
	for _, sample := range samples {
		if uint64(sample) < min {
			min = uint64(sample)
		}
	}

	return time.Duration(min)
}

// GetMaximum return the maximum value from all samples in the timer.
func (t *Timer) GetMaximum() time.Duration {
	samples := t.GetAll()
	if len(samples) == 0 {
		return 0
	}

	max := uint64(0)
	for _, sample := range samples {
		if uint64(sample) > max {
			max = uint64(sample)
		}
	}

	return time.Duration(max)
}
