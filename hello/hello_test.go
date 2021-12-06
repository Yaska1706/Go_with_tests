package hello

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Yaska")
	want := "Hello, Yaska"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
