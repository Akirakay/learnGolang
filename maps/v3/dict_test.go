package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dict{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got '%s' want '%s' given, '%s'", got, want, "test")
	}

}
