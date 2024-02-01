package problem_7

import "testing"

var tests = []struct {
	name     string
	input    uint
	expected uint64
}{
	{
		name:     "First Prime",
		input:    1,
		expected: 2,
	},
	{
		name:     "Second Prime",
		input:    2,
		expected: 3,
	},
	{
		name:     "Sixth Prime",
		input:    6,
		expected: 13,
	},
}

func TestNthPrime(t *testing.T) {
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			res := nthPrime(tt.input)
			if res != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, res)
			}
		})
	}
}
