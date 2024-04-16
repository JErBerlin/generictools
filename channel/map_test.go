package channel

import (
	"math"
	"sync"
	"testing"
)

const (
	tol = 1e-12
)

// TestMapTypeFloat64s tests the Map function with float64 values.
func TestMapTypeFloat64s(t *testing.T) {
	in := make(chan float64, 3)
	out := make(chan float64, 3)

	floats := []float64{1.1, 2.2, 3.3}
	want := []float64{2.2, 3.3, 4.4}

	// Fill the input channel in a separate goroutine
	go func() {
		for _, f := range floats {
			in <- f
		}
		close(in)
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	// Process the mapping in a separate goroutine
	go func() {
		defer wg.Done()
		Map(in, out, func(f float64) float64 { return f + 1.1 })
	}()

	wg.Wait()
	close(out)

	// Collect the output from the out channel
	var got []float64
	for f := range out {
		got = append(got, f)
	}

	// Compare the result
	if !equalsFloat64Slice(got, want, tol) {
		t.Errorf("MapChan(float64s) = %v, want %v", got, want)
	}
}

// equalsFloat64Slice compares two slices of float64s for equality within a given tolerance.
func equalsFloat64Slice(a, b []float64, tol float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Abs(a[i]-b[i]) > tol {
			return false
		}
	}
	return true
}
