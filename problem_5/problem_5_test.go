package problem_5

import "testing"

func TestDivisibleByRange(t *testing.T) {
	want := int64(2520)

	got := divisibleByRange(1, 10)

	if got != want {
		t.Errorf("divisibleByRange(1, 10) = %d; want %d", got, want)
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve()
	}
}
