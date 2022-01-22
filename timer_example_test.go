package stats_test

import (
	"fmt"
	"time"

	"github.com/ecnepsnai/stats"
)

func ExampleTimer_AddSample() {
	timer := stats.NewTimer(5)
	timer.AddSample(5 * time.Minute)
	fmt.Printf("Last sample: %d minutes", int(timer.GetLast().Minutes()))
	// Output: Last sample: 5 minutes
}

func ExampleTimer_Reset() {
	timer := stats.NewTimer(5)
	timer.AddSample(5 * time.Minute)
	timer.Reset()
	fmt.Printf("Last sample: %d minutes", int(timer.GetLast().Minutes()))
	// Output: Last sample: 0 minutes
}

func ExampleTimer_GetLast() {
	timer := stats.NewTimer(5)
	timer.AddSample(5 * time.Minute)
	fmt.Printf("Last sample: %d minutes", int(timer.GetLast().Minutes()))
	// Output: Last sample: 5 minutes
}

func ExampleTimer_GetAll() {
	timer := stats.NewTimer(5)
	timer.AddSample(5 * time.Minute)
	timer.AddSample(3 * time.Minute)
	fmt.Printf("All samples: %v", timer.GetAll())
	// Output: All samples: [5m0s 3m0s]
}

func ExampleTimer_GetAverage() {
	timer := stats.NewTimer(5)
	timer.AddSample(5 * time.Minute)
	timer.AddSample(3 * time.Minute)
	fmt.Printf("Average: %d minutes", int(timer.GetAverage().Minutes()))
	// Output: Average: 4 minutes
}

func ExampleTimer_GetMinimum() {
	timer := stats.NewTimer(5)
	timer.AddSample(5 * time.Minute)
	timer.AddSample(3 * time.Minute)
	fmt.Printf("Minimum: %d minutes", int(timer.GetMinimum().Minutes()))
	// Output: Minimum: 3 minutes
}

func ExampleTimer_GetMaximum() {
	timer := stats.NewTimer(5)
	timer.AddSample(5 * time.Minute)
	timer.AddSample(3 * time.Minute)
	fmt.Printf("Maximum: %d minutes", int(timer.GetMaximum().Minutes()))
	// Output: Maximum: 5 minutes
}
