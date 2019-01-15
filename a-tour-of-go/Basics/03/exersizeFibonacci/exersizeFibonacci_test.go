package exersizeFibonacci

import (
	"testing"
)

func TestFibFunc(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{
			0,
			0,
		},
		{
			1,
			1,
		},
		{
			2,
			1,
		},
		{
			3,
			2,
		},
		{
			4,
			3,
		},
		{
			5,
			5,
		},
		{
			6,
			8,
		},
	}

	for i, test := range tests {
		if fib(test.input) != test.expected {
			t.Fatalf("fib(%d) returned wrong value. got=%d, expected=%d", i, fib(test.input), test.expected)
		}
	}
}
