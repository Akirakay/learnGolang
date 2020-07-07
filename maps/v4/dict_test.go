package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dict{"test": "this is just a test"}

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertStrings(t, err, ErrNotFound)
	})
}

func assertStrings(t *testing.T, got, want error) {
	if got != want {
		t.Errorf("got '%s' want '%s' given, '%s'", got, want, "test")
	}

}
