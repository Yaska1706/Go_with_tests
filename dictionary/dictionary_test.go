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
	t.Run("new word", func(t *testing.T) {

		dictionary := Dictionary{}
		word := "practice"
		definition := "makes perfect"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})

}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	want := definition
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word", err)
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {

		word := "practice"
		definition := "makes perfect"
		newdefinition := "more improvement"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, newdefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newdefinition)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test2"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordNotFound)
	})

}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "another test"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrDoesNotExist {
		t.Errorf("Expected %q to be deleted", word)
	}
}
