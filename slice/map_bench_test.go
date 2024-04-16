package slice

import (
	"runtime"
	"strings"
	"testing"
)

// generateIntSlice generates a slice of int of size n for benchmarking
func generateIntSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i
	}
	return s
}

// generateFloat64Slice generates a slice of float64 of size n for benchmarking.
func generateFloat64Slice(n int) []float64 {
	s := make([]float64, n)
	for i := 0; i < n; i++ {
		s[i] = float64(i) * 6.28318
	}
	return s
}

// generateStringSlice generates a slice of strings of size n for benchmarking.
func generateStringSlice(n int) []string {
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] += "one more step"
	}
	return s
}

// BenchmarkGenericMapTypeInt benchmarks the generic Map function with type int slices.
func BenchmarkGenericMapTypeInt(b *testing.B) {
	s := generateIntSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Map(s, func(n int) int { return n + 1 })
	}
}

// BenchmarkOldSchoolMapInts benchmarks the MapInts function.
func BenchmarkOldSchoolMapInts(b *testing.B) {
	s := generateIntSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MapInts(s, func(n int) int { return n + 1 })
	}
}

// BenchmarkGenericMapFloat64 benchmarks the generic Map function with type float64 slices.
func BenchmarkGenericMapTypeFloat64(b *testing.B) {
	s := generateFloat64Slice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Map(s, func(f float64) float64 { return f + 1.1 })
	}
}

// BenchmarkOldSchoolMapFloat64s benchmarks the MapFloat64s function.
func BenchmarkOldSchoolMapFloat64s(b *testing.B) {
	s := generateFloat64Slice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MapFloat64s(s, func(f float64) float64 { return f + 1.1 })
	}
}

// BenchmarkGenericMapTypeString benchmarks the generic Map function with type string slices.
func BenchmarkGenericMapTypeString(b *testing.B) {
	s := generateStringSlice(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Map(s, func(str string) string { return strings.ToUpper(str) })
	}
}

// BenchmarkOldSchoolMapStrings benchmarks the MapStrings function.
func BenchmarkOldSchoolMapStrings(b *testing.B) {
	s := generateStringSlice(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MapStrings(s, func(str string) string { return strings.ToUpper(str) })
	}
}

// BenchmarkGenericMapTypeStringParallel benchmarks the parallelized generic Map function with type string slices.
func BenchmarkGenericMapTypeStringParallel(b *testing.B) {
	b.Logf("number of CPUs: %d\n", runtime.NumCPU())
	s := generateStringSlice(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MapParallel(s, func(str string) string { return strings.ToUpper(str) })
	}
}

// BenchmarkOldSchoolMapStringsParallel benchmarks the parallelized MapStrings function.
func BenchmarkOldSchoolMapStringsParallel(b *testing.B) {
	b.Logf("number of CPUs: %d\n", runtime.NumCPU())
	s := generateStringSlice(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MapStringsParallel(s, func(str string) string { return strings.ToUpper(str) })
	}
}
