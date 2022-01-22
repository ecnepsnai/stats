package stats_test

import (
	"testing"

	"github.com/ecnepsnai/stats"
)

func TestCounter(t *testing.T) {
	c := stats.NewCounter()
	c.Increment()
	if c.Get() != 1 {
		t.Errorf("Unexpected counter value. Expected 1 got %d", c.Get())
	}
	c.IncrementBy(2)
	if c.Get() != 3 {
		t.Errorf("Unexpected counter value. Expected 3 got %d", c.Get())
	}
	c.Decrement()
	if c.Get() != 2 {
		t.Errorf("Unexpected counter value. Expected 2 got %d", c.Get())
	}
	c.DecrementBy(2)
	if c.Get() != 0 {
		t.Errorf("Unexpected counter value. Expected 0 got %d", c.Get())
	}
	c.Set(5)
	if c.Get() != 5 {
		t.Errorf("Unexpected counter value. Expected 5 got %d", c.Get())
	}
	c.DecrementBy(15)
	if c.Get() != 0 {
		t.Errorf("Unexpected counter value. Expected 0 got %d", c.Get())
	}
}
