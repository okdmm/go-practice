package functions

import "testing"

func TestAddFunc(t *testing.T) {
	tests := []struct {
		input    [2]int
		expected int
	}{
		{
			[2]int{13, 42},
			55,
		},
		{
			[2]int{1, 2},
			3,
		},
		{
			[2]int{-6, 10},
			4,
		},
	}

	for _, test := range tests {
		expected := add(test.input[0], test.input[1])
		if expected != test.expected {
			t.Fatalf("add(%d, %d) is wrong value.expected=%d, got=%d", test.input[0], test.input[1], test.expected, expected)
		}
	}
}
