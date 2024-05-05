package arrayslice

import "testing"

func TestSum(t *testing.T) {
	t.Run("Collection of non fixed size numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		actual := Sum(numbers)
		expected := 6

		if actual != expected {
			t.Errorf("got %d want %d given, %v", actual, expected, numbers)
		}
	})
}
