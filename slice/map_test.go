package slice

import (
	"math"
	"reflect"
	"testing"
)

const (
	tol = 1e-12
)

// Test type specific Map functions

func TestMapInts(t *testing.T) {
	ints := []int{1, 2, 3}
	want := []int{2, 3, 4}
	got := MapInts(ints, func(n int) int { return n + 1 })

	if !reflect.DeepEqual(got, want) {
		t.Errorf("MapInts(ints) = %v, want %v", got, want)
	}
}

func TestMapFloat64s(t *testing.T) {
	floats := []float64{1.1, 2.2, 3.3}
	want := []float64{2.2, 3.3, 4.4}
	got := MapFloat64s(floats, func(f float64) float64 { return f + 1.1 })

	if !equalsFloat64Slice(got, want, tol) {
		t.Errorf("MapFloat64s(floats) = %v, want %v", got, want)
	}

}

func TestMapRunes(t *testing.T) {
	runes := []rune{'a', 'b', 'c'}
	want := []rune{'b', 'c', 'd'}
	got := MapRunes(runes, func(r rune) rune { return r + 1 })

	if !reflect.DeepEqual(got, want) {
		t.Errorf("MapRunes(runes) = %v, want %v", got, want)
	}
}

func TestMapStrings(t *testing.T) {
	strings := []string{"a", "b", "c"}
	want := []string{"aa", "bb", "cc"}
	got := MapStrings(strings, func(s string) string { return s + s })

	if !reflect.DeepEqual(got, want) {
		t.Errorf("MapStrings(strings) = %v, want %v", got, want)
	}
}

func TestMapTypeInt(t *testing.T) {
	ints := []int{1, 2, 3}
	want := []int{2, 3, 4}
	got := Map(ints, func(n int) int { return n + 1 })

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Map(ints) = %v, want %v", got, want)
	}
}

// Test generic types function Map

func TestMapTypeFloat64s(t *testing.T) {
	floats := []float64{1.1, 2.2, 3.3}
	want := []float64{2.2, 3.3, 4.4}
	got := Map(floats, func(f float64) float64 { return f + 1.1 })

	if !equalsFloat64Slice(got, want, tol) {
		t.Errorf("Map(floats) = %v, want %v", got, want)
	}

}

func TestMapTypeRune(t *testing.T) {
	runes := []rune{'a', 'b', 'c'}
	want := []rune{'b', 'c', 'd'}
	got := Map(runes, func(r rune) rune { return r + 1 })

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Map(runes) = %v, want %v", got, want)
	}
}

func TestMapTypeString(t *testing.T) {
	strings := []string{"a", "b", "c"}
	want := []string{"aa", "bb", "cc"}
	got := Map(strings, func(s string) string { return s + s })

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Map(strings) = %v, want %v", got, want)
	}
}

// equalsFloat64Slice is a help function that compares two slices of float64s for equality within a given tolerance.
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
