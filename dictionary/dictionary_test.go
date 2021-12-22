package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {

	dictionary := Dictionary{
		"test": "this is a test",
	}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"

		assertStrings(t, got, want)
	})

	t.Run("Unkown word ", func(t *testing.T) {
		_, got := dictionary.Search("new")

		if got == nil {
			t.Fatal("expected error")
		}
		assertError(t, got, ErrDoesNotExist)
	})

}

func assertStrings(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want error %q", got, want)
	}
}

func TestDictionary_Add(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("practice", "makes perfect")

	want := "makes perfect"
	got, err := dictionary.Search("practice")
	if err != nil {
		t.Fatal("should find added word", err)
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
