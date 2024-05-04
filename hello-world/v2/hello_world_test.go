package main

import "testing"

func TestHelloWorld(t *testing.T) {
	actual := Hello()
	expected := "Hello, world"

	if actual != expected {
		t.Errorf("Actual: %q Expected: %q", actual, expected)
	}
}
