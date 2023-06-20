package binarytree

import (
	"testing"
)

func TestMaxDepth(t *testing.T) {
	root := &Node{val: 3}
	root.Left = &Node{val: 9}
	root.Right = &Node{val: 20}
	root.Right.Left = &Node{val: 15}
	root.Right.Right = &Node{val: 7}

	t.Run("Test full", func(t *testing.T) {
		want := 3

		got := maxDepth(root)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test empty tree", func(t *testing.T) {
		want := 0

		got := maxDepth(nil)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
