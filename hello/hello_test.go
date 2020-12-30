package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	}

	t.Run("Hello world", func(t *testing.T) {
		got := Hello("Amogh", "english")
		want := "Hello Amogh"
		assertCorrectMessage(t, got, want)

	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello World"

		// if got != want {
		// 	t.Errorf("got %q want %q", got, want)
		// }

		assertCorrectMessage(t, got, want)
	})

	t.Run("Hello world in spanish", func(t *testing.T) {

		got := Hello("Elodie", "Spanish")
		want := "Holla Elodie"
		assertCorrectMessage(t, got, want)

	})

	t.Run("Hello world in French", func(t *testing.T) {

		got := Hello("Elodie", "French")
		want := "Bonjour Elodie"
		assertCorrectMessage(t, got, want)

	})

}
