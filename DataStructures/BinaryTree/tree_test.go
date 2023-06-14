package binarytree

import (
	"reflect"
	"testing"
)

func TestTreeBuildFromInAndPostOrderLists(t *testing.T) {
	// Lits to build from
	InOrder := []int{4, 8, 2, 5, 1, 6, 3, 7}
	PostOrder := []int{8, 4, 5, 2, 6, 7, 3, 1}

	// Root
	root := new(BinaryTree[int])
	root.BuildFromInAndPostOrderLists(InOrder, PostOrder)

	t.Run("Assert Size of Array", func(t *testing.T) {
		want := 8

		got := root.Size()

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Assert Preorder", func(t *testing.T) {
		want := []int{1, 2, 4, 8, 5, 3, 6, 7}

		got := root.WalkPreOrder()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
