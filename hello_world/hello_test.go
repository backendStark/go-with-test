package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Dmitrii")
	want := "Hello, Dmitrii"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
