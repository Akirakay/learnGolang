package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := map[string]string{"test": "this is just a test"}

	got := Search(dict, "test")
	want := "this is just a test"

	assertStrings(t, got, want)
}

func assertStrings(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got '%s' want '%s' given, '%s'", got, want, "test")
	}
}
