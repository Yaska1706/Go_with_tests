package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	arr := [5]int{2, 3, 4, 5, 1}

	got := Sum(arr[:])
	want := 15

	if got != want {
		t.Errorf("got %d want %d ", got, want)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{3, 5}, []int{3, 5, 2})
	want := []int{8, 10}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d ", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checksums := func(t *testing.T, got, want int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("with 2 slices", func(t *testing.T) {
		got := SumAllTails([]int{3, 4}, []int{7, 6})
		want := 10
		checksums(t, got, want)

	})
	t.Run("with 3 slices", func(t *testing.T) {
		got := SumAllTails([]int{3, 4}, []int{7, 6}, []int{3, 4})
		want := 14

		checksums(t, got, want)

	})
	t.Run("with some empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{7, 6}, []int{})
		want := 6

		checksums(t, got, want)

	})
}
