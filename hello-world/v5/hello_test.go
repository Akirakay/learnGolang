package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello from people", func(t *testing.T) {
		got := Hello("Akira")
		want := "hello golang -- from Akira"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello from when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "hello golang -- from Akira"

		assertCorrectMessage(t, got, want)
	})
}
