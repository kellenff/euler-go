package problem_3

import (
	"reflect"
	"testing"
)

func TestGetPrimeFactors(t *testing.T) {
	want := []uint64{5, 7, 13, 29}

	got := getPrimeFactors(13195)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("getPrimeFactors(13195) = %d, want %d", got, want)
	}
}
