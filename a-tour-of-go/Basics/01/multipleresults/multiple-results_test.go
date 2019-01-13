package multipleresults

import "testing"

func TestSwapFunc(t *testing.T) {
	tests := []struct {
		input    [2]string
		expected [2]string
	}{
		{
			[2]string{"hello", "world"},
			[2]string{"world", "hello"},
		},
		{
			[2]string{" ", "a"},
			[2]string{"a", " "},
		},
	}

	for _, test := range tests {
		a, b := swap(test.input[0], test.input[1])
		if a != test.expected[0] || b != test.expected[1] {
			t.Fatalf("not swaped. expected = %s, %s. got=%s, %s", test.expected[0], test.expected[1], a, b)
		}
	}
}
