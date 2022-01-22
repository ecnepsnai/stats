package stats_test

import (
	"testing"
	"time"

	"github.com/ecnepsnai/stats"
)

func TestTimedCounter(t *testing.T) {
	tc := stats.NewTimedCounter(2 * time.Second)

	tc.Increment()
	tc.Increment()
	time.Sleep(1100 * time.Millisecond)
	tc.Increment()
	if last := tc.GetLast(1 * time.Second); last != 1 {
		t.Errorf("Unexpected last value from timed counter. Expected 1 got %d", last)
	}
	if all := tc.GetAll(); all != 3 {
		t.Errorf("Unexpected all value from timed counter. Expected 3 got %d", all)
	}
	tc.DecrementLast()
	if all := tc.GetAll(); all != 2 {
		t.Errorf("Unexpected all value from timed counter. Expected 2 got %d", all)
	}
	time.Sleep(2 * time.Second)
	tc.Cleanup()
	if all := tc.GetAll(); all != 0 {
		t.Errorf("Unexpected all value from timed counter. Expected 0 got %d", all)
	}
	tc.Increment()
	tc.Cleanup()
	if all := tc.GetAll(); all != 1 {
		t.Errorf("Unexpected all value from timed counter. Expected 1 got %d", all)
	}
	tc.DecrementLast()
	if all := tc.GetAll(); all != 0 {
		t.Errorf("Unexpected all value from timed counter. Expected 0 got %d", all)
	}
	tc.DecrementLast()
	tc.Cleanup()
}
