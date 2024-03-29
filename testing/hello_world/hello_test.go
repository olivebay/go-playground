package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T){
		got := Hello("Chris", "english")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to World", func(t *testing.T){
		got := Hello("", "english")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in spanish", func(t *testing.T){
		got := Hello("Chris", "Spanish")
		want := "Hola, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in french", func(t *testing.T){
		got := Hello("Melina", "french")
		want := "Bonjour, Melina"
		assertCorrectMessage(t, got, want)
	})

}