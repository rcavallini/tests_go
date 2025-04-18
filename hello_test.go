package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Renan")
	want := "Hello, Renan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

/*
func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
*/
