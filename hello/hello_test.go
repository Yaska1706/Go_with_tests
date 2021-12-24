package hello

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Yaska", "")
		want := "Hello, Yaska"

		assertCorrectMessage(t, got, want)
	})
	t.Run("use 'World' when string is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Yaska", "Spanish")
		want := "Hola, Yaska"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Yaska", "French")
		want := "Bonjour, Yaska"

		assertCorrectMessage(t, got, want)
	})

}

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Yaska")

	got := buffer.String()
	want := "Hello, Yaska"

	if got != want {
		t.Errorf("expected %q got %q", want, got)
	}
}
