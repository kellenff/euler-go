package problem_10

import (
	"testing"
)

func TestSumOfPrimes(t *testing.T) {
	tests := []struct {
		name  string
		limit uint64
		want  uint64
	}{
		{name: "PrimeSumUnder10", limit: 10, want: 17},
		{name: "PrimeSumUnder2", limit: 2, want: 0},
		{name: "PrimeSumUnder0", limit: 0, want: 0},
		{name: "PrimeSumUnder5", limit: 5, want: 5},
		{name: "PrimeSumUnder20", limit: 20, want: 77},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfPrimes(tt.limit); got != tt.want {
				t.Errorf("sumOfPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
