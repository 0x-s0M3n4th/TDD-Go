package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Somenath") // variables
	want := "Hello, Somenath!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
