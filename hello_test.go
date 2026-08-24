package main

import "testing"

func TestHello(t *testing.T) {
	// Case when argument is provided into the Hello() func
	// Spanish

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Somenath", "Spanish")
		want := "Hola, Somenath:)"
		assertCorrectMessage(t, got, want)
	})

	// French

	t.Run("in french", func(t *testing.T) {
		got := Hello("Levi", "French")
		want := "Bonjour, Levi^-^"
		assertCorrectMessage(t, got, want)
	})

	// Case when an empty string  is provided as argument
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // This tells the testing framework that  if any error occurs don't point it towards me, i am just a helper function.
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
