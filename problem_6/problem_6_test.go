package problem_6

import "testing"

func TestSumOfSquares(t *testing.T) {
	want := int64(385)

	got := sumOfSquares(1, 11)

	if want != got {
		t.Fatalf("sumOfSquares() = %d, want %d", got, want)
	}
}

func TestSquareOfSum(t *testing.T) {
	want := int64(3025)

	got := squareOfSum(1, 11)

	if want != got {
		t.Fatalf("squareOfSum() = %d, %d", got, want)
	}
}
