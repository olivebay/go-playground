package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	word := "test"
	definition := "just a test"
	dictionary := Dictionary{word: definition}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search(word)
		want := "just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("add new word", func(t *testing.T) {
		word := "test"
		definition := "just a test"
		dictionary := Dictionary{}

		got := dictionary.Add(word, definition)

		assertError(t, got, nil)
	})

	t.Run("add existig word", func(t *testing.T) {
		word := "test"
		definition := "just a test"
		dictionary := Dictionary{word: definition}

		got := dictionary.Add(word, definition)

		assertError(t, got, ErrWordsExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "just a test"
		dictionary := Dictionary{}

		got := dictionary.Add(word, definition)

		assertError(t, got, nil)
	})

	t.Run("existig word", func(t *testing.T) {
		word := "test"
		definition := "just a test"
		newDefinition := "new word"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

}

func (d Dictionary) TestDelete(t *testing.T) {
	word := "test"
	definition := "just a test"
	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("expected  %q to be deleted", word)
	}

}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}

}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %w", got, definition)
	}
}
