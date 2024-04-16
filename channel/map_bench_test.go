package channel

import (
	"strings"
	"sync"
	"testing"
)

// generateStringSlice generates a slice of strings of size n for benchmarking.
func generateStringSlice(n int) []string {
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] += "one more step"
	}
	return s
}

func BenchmarkMapTypeString(b *testing.B) {
	s := generateStringSlice(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		in := make(chan string, len(s))
		out := make(chan string, len(s))
		var wg sync.WaitGroup

		// Start the Map function in a goroutine.
		wg.Add(1)
		go func() {
			defer wg.Done()
			Map(in, out, strings.ToUpper)
		}()

		// Send input strings to the in channel in another goroutine.
		go func() {
			for _, str := range s {
				in <- str
			}
			close(in)
		}()

		// Wait for the Map function to process all inputs.
		wg.Wait()
		close(out) // Now it's safe to close the out channel.

		// Drain the out channel.
		for range out {
		}
	}
}

func BenchmarkMapTypeStringParallel(b *testing.B) {
	s := generateStringSlice(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		in := make(chan string, len(s))
		out := make(chan string, len(s))
		var wg sync.WaitGroup

		// Start the Map function in a goroutine.
		wg.Add(1)
		go func() {
			defer wg.Done()
			MapParallel(in, out, strings.ToUpper)
		}()

		// Send input strings to the in channel in another goroutine.
		go func() {
			for _, str := range s {
				in <- str
			}
			close(in)
		}()

		// Wait for the Map function to process all inputs.
		wg.Wait()
		close(out) // Now it's safe to close the out channel.

		// Drain the out channel.
		for range out {
		}
	}
}
