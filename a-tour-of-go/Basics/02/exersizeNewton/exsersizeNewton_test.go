package exersizeNewton

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	input := 2.0
	expectedDiff := 1e-15

	v := Sqrt(input)
	gotDiff := math.Abs(input - v*v)
	if gotDiff > expectedDiff {
		t.Fatalf("much Diff %v", gotDiff)
	}
}
