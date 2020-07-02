package main

import "testing"

func TestHello(t *testing.T) {
	got := hello()
	want := "hello golang -- from Akira"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
