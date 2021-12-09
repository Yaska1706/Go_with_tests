package stacks

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("Push", func(t *testing.T) {
		stack := Stack{}
		got := stack.Push(ItemType{2, 2, 1})
		want := []ItemType{2, 2, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
