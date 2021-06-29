package ssort

import (
	"testing"
	"sort"
	"math/rand"
)

func randSlice(size int) []float64 {
	out := make([]float64, size)
	for i := 0; i < size; i++ {
		out[i] = rand.Float64()
	}
	return out
}

func TestFloat64sSimple(t *testing.T) {
	cases := []struct {
		name string
		slice []float64
	}{
		{"10", randSlice(10)},
		{"1,000", randSlice(1000)},
		{"100,000", randSlice(100000)},
		{"10,000,000", randSlice(10000000)},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			Float64s(c.slice)
			if !sort.Float64sAreSorted(c.slice) {
				t.Errorf("Slice is not sorted in ascending order")
			}
		})
	}
}

func TestFloat64sNaN(t *testing.T) {
}

func TestFloat64sPlusInf(t *testing.T) {
}

func TestFloat64sMinusInf(t *testing.T) {
}

func TestFloat64sMixed(t *testing.T) {
}
