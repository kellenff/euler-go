package problem_2

import (
	"reflect"
	"testing"
)

func TestGetFibs(t *testing.T) {
	want := []uint64{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	got := getFibs(90)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("getFibs() = %v, want %v", got, want)
	}
}
