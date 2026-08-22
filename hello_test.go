package main

import "testing"

func TestHello(t *testing.T) {
	// Case when argument is provided into the Hello() func
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Somenath")
		want := "Hello, Somenath!"
		assertCorrectMessage(t, got, want)
	})

	// Case when an empty string  is provided as argument
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
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
