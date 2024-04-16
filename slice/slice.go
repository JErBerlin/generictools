package slice

import (
	"runtime"
	"sync"
)

func Map[T any](s []T, f func(T) T) []T {
	r := make([]T, len(s))

	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

func MapInts(s []int, f func(int) int) []int {
	r := make([]int, len(s))

	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

func MapFloat64s(s []float64, f func(float64) float64) []float64 {
	r := make([]float64, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

func MapRunes(s []rune, f func(rune) rune) []rune {
	r := make([]rune, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

func MapStrings(s []string, f func(string) string) []string {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = f(e)
	}
	return r
}

// MapParallel applies the function f to each element of slice s in parallel by segmenting the slice.
// This implementation is safe for concurrent use because each goroutine writes to a unique index
// of the pre-allocated slice, avoiding concurrent modifications to the slice header and
// since no resize or reallocation is happening.
func MapParallel[T any](s []T, f func(T) T) []T {
	r := make([]T, len(s))
	var wg sync.WaitGroup

	// Determine the number of goroutines to use.
	// divide the slice into a number of segments corresponding to the number of logical CPUs.
	numGoroutines := runtime.NumCPU()
	segmentSize := len(s) / numGoroutines
	if segmentSize == 0 {
		segmentSize = 1 // Ensure at least one element per segment if slice size < numGoroutines.
	}

	for i := 0; i < numGoroutines; i++ {
		// start and end indices for the segment.
		start := i * segmentSize
		end := start + segmentSize
		if i == numGoroutines-1 || end > len(s) {
			end = len(s) - 1 // Ensure the last segment goes to the end of the slice.
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for j := start; j <= end; j++ {
				r[j] = f(s[j])
			}
		}(start, end)
	}

	wg.Wait()
	return r
}

// MapStringsParallel applies the function f to each element of slice s in parallel.
// This implementation is safe for concurrent use because each goroutine writes to a unique index
// of the pre-allocated slice, avoiding concurrent modifications to the slice header and
// since no resize or reallocation is happening.
func MapStringsParallel(s []string, f func(string) string) []string {
	r := make([]string, len(s))
	var wg sync.WaitGroup

	// Determine the number of goroutines to use.
	// divide the slice into a number of segments corresponding to the number of logical CPUs.
	numGoroutines := runtime.NumCPU()
	segmentSize := len(s) / numGoroutines
	if segmentSize == 0 {
		segmentSize = 1 // Ensure at least one element per segment if slice size < numGoroutines.
	}

	for i := 0; i < numGoroutines; i++ {
		// start and end indices for the segment.
		start := i * segmentSize
		end := start + segmentSize
		if i == numGoroutines-1 || end > len(s) {
			end = len(s) - 1 // Ensure the last segment goes to the end of the slice.
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for j := start; j <= end; j++ {
				r[j] = f(s[j])
			}
		}(start, end)
	}

	wg.Wait()
	return r
}
