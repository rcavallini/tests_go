package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got, err := Hello("Chris", "English")
		want := "Hello, Chris"
		assertNoError(t, err)
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got, err := Hello("", "English")
		want := "Hello, World"
		assertNoError(t, err)
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got, err := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertNoError(t, err)
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("did not expect error but got one: %v", err)
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
