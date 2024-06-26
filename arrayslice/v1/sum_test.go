package arrayslice

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	actual := Sum(numbers)
	expected := 15

	if actual != expected {
		t.Errorf("got %d want %d given, %v", actual, expected, numbers)
	}
}
