package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Jay", "")
		want := "Hello, Jay"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Gus", spanish)
		want := "Hola, Gus"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", french)
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in German", func(t *testing.T) {
		got := Hello("Greta", german)
		want := "Hallo, Greta"
		assertCorrectMessage(t, got, want)
	})
}

// For helper functions, it's a good idea to accept a testing.TB which
// is an interface that *testing.T and *testing.B both satisfy, so you can
// call helper functions from a test, or a benchmark
func assertCorrectMessage(t testing.TB, got, want string) {

	// t.Helper() is needed to tell the test suite that this method is a helper.
	// By doing this, when it fails, the line number reported will be in our function
	// call rather than inside our test helper.
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
