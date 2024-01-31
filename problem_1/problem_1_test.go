package problem_1

import "testing"

func TestFindSum(t *testing.T) {
	got := Solve(10)
	want := uint(23)

	if got != want {
		t.Fatalf("Got %d, want %d", got, want)
	}
}
