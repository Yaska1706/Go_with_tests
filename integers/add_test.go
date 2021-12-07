package integers

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(10, 10)
	want := 20

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestModulus(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Is divisible by 2", func(t *testing.T) {
		got := Modulus("10")
		want := "Divisible by 2"

		assertCorrectMessage(t, got, want)
	})

}
