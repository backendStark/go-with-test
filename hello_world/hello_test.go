package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("should type the name", func(t *testing.T) {
		got := Hello("Dmitrii")
		want := "Hello, Dmitrii"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("should type 'Hello, world' with epmty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
