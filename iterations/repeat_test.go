package iterations

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
