package stats_test

import (
	"testing"
	"time"

	"github.com/ecnepsnai/stats"
)

func TestTimer(t *testing.T) {
	timer := stats.NewTimer(3)
	timer.AddSample(5 * time.Minute)
	timer.AddSample(2 * time.Minute)
	timer.AddSample(210 * time.Second)

	if all := timer.GetAll(); len(all) != 3 {
		t.Errorf("Unexpected numbre of values in timer. Expected 3 got %d", len(all))
	}

	if avg := timer.GetAverage(); avg.Seconds() != 210 {
		t.Errorf("Unexpected average value from timer. Expected 210 got %f", avg.Seconds())
	}

	if last := timer.GetLast(); last.Minutes() != 3.5 {
		t.Errorf("Unexpected last value from timer. Expected 3.5 got %f", last.Minutes())
	}

	if max := timer.GetMaximum(); max.Minutes() != 5 {
		t.Errorf("Unexpected max value from timer. Expected 5 got %f", max.Minutes())
	}

	if min := timer.GetMinimum(); min.Minutes() != 2 {
		t.Errorf("Unexpected min value from timer. Expected 2 got %f", min.Minutes())
	}

	timer.Reset()
	if all := timer.GetAll(); len(all) != 0 {
		t.Errorf("Unexpected numbre of values in timer. Expected 0 got %d", len(all))
	}

	if avg := timer.GetAverage(); avg.Seconds() != 0 {
		t.Errorf("Unexpected average value from timer. Expected 0 got %f", avg.Seconds())
	}

	if last := timer.GetLast(); last.Minutes() != 0 {
		t.Errorf("Unexpected last value from timer. Expected 0 got %f", last.Minutes())
	}

	if max := timer.GetMaximum(); max.Minutes() != 0 {
		t.Errorf("Unexpected max value from timer. Expected 0 got %f", max.Minutes())
	}

	if min := timer.GetMinimum(); min.Minutes() != 0 {
		t.Errorf("Unexpected min value from timer. Expected 0 got %f", min.Minutes())
	}
}
