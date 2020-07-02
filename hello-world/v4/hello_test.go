package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello from people", func(t *testing.T) {
		got := Hello("Akira")
		want := "hello golang -- from Akira"

		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})

	t.Run("say hello from when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "hello golang -- from Akira"

		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})
}
