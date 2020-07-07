package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "akira")

	got := buffer.String()
	want := "Hello, akira"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
