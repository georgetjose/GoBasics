package arrayslice

import "testing"

func TestSum(t *testing.T) {
	t.Run("Collection of non fixed size slices", func(t *testing.T) {
		actualSlice := SumAll([]int{1, 2}, []int{0, 9})
		expectedSlice := []int{3, 9}

		for i := 0; i < len(actualSlice); i++ {
			if actualSlice[i] != expectedSlice[i] {
				t.Errorf("Actual: %d Expected: %d ", actualSlice, expectedSlice)
			}
		}

	})
}
