package integer

import "testing"

func TestAdd(t *testing.T) {
	actual := Add(2, 5)
	expected := 7

	if actual != expected {
		t.Errorf("Expected: %q Actual: %q", expected, actual)
	}
}
