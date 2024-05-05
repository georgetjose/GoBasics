package arrayslice

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of non fixed size slices", func(t *testing.T) {
		actualSlice := SumAll([]int{1, 2}, []int{0, 9})
		expectedSlice := []int{3, 9}

		if !reflect.DeepEqual(actualSlice, expectedSlice) {
			t.Errorf("Actual: %d Expected: %d ", actualSlice, expectedSlice)
		}
	})

	t.Run("Collection of non fixed size slices", func(t *testing.T) {
		actualSlice := SumAll([]int{1, 2, 5})
		expectedSlice := []int{8}

		if !slices.Equal(actualSlice, expectedSlice) {
			t.Errorf("Actual: %d Expected: %d ", actualSlice, expectedSlice)
		}
	})
}
