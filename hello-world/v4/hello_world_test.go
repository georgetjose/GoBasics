package main

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Run("Saying Hello with name", func(t *testing.T) {
		actual := Hello("George")
		expected := "Hello, George"
		assertCorrectMessage(actual, expected, t)
	})
	t.Run("saying Hello world when name is not provided", func(t *testing.T) {
		actual := Hello("")
		expected := "Hello, world"
		assertCorrectMessage(actual, expected, t)
	})
}

//testing.TB which is an interface that *testing.T and *testing.B both satisfy, so you can call helper functions from a test
func assertCorrectMessage(actual, expected string, t testing.TB) {
	t.Helper()//t.Helper() is needed to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
	if actual != expected {
		t.Errorf("Actual: %q Expected: %q", actual, expected)
	}
}
