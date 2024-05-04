package main

import "testing"

func TestHelloWorld(t *testing.T) {
	actual := Hello("George")
	expected := "Hello, George"

	if actual != expected {
		t.Errorf("Actual: %q Expected: %q", actual, expected)
	}
}
