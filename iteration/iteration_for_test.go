package iteration

import (
	"fmt"
	"testing"
)

func TestIterationFor(t *testing.T) {
	actual := repeat("b",8)
	expected := "bbbbbbbb"

	if actual != expected {
		t.Errorf("Actual: %q Expected: %q", actual, expected)
	}
}

func BenchMarkIterationFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(repeat("a",8))
	}
}
