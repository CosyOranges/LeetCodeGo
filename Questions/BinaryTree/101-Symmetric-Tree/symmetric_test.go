package binarytree

import (
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 3}

	t.Run("Test for Symmetrical input", func(t *testing.T) {
		if !isSymmetric(root) {
			t.Errorf("got: false, want: true")
		}
	})

	t.Run("Test for Assymetrical Tree", func(t *testing.T) {
		root.Left.Left.Val = 3

		if isSymmetric(root) {
			t.Errorf("got: true, want: false")
		}
	})
}
