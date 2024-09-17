package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jay")
	want := "Hello, Jay"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
