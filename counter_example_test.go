package stats_test

import (
	"fmt"

	"github.com/ecnepsnai/stats"
)

func ExampleCounter_Increment() {
	counter := stats.NewCounter()
	counter.Increment()
	fmt.Printf("Current value: %d", counter.Get())
	// Output: Current value: 1
}

func ExampleCounter_Decrement() {
	counter := stats.NewCounter()
	counter.Set(6)
	counter.Decrement()
	fmt.Printf("Current value: %d", counter.Get())
	// Output: Current value: 5
}

func ExampleCounter_IncrementBy() {
	counter := stats.NewCounter()
	counter.IncrementBy(5)
	fmt.Printf("Current value: %d", counter.Get())
	// Output: Current value: 5
}

func ExampleCounter_DecrementBy() {
	counter := stats.NewCounter()
	counter.Set(7)
	counter.DecrementBy(5)
	fmt.Printf("Current value: %d", counter.Get())
	// Output: Current value: 2
}

func ExampleCounter_Set() {
	counter := stats.NewCounter()
	counter.Set(8)
	fmt.Printf("Current value: %d", counter.Get())
	// Output: Current value: 8
}

func ExampleCounter_Get() {
	counter := stats.NewCounter()
	counter.Set(8)
	fmt.Printf("Current value: %d", counter.Get())
	// Output: Current value: 8
}
