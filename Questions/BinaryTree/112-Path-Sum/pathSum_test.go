package binarytree

import "testing"

func TestHasPathSum(t *testing.T) {
	t.Run("Test negatives", func(t *testing.T) {
		tree := &TreeNode{Val: -2}
		tree.Right = &TreeNode{Val: -3}

		if !hasPathSum(tree, -5) {
			t.Errorf("got: false, want: true")
		}
	})

	t.Run("Test larger tree", func(t *testing.T) {
		tree := &TreeNode{Val: 5}
		tree.Left = &TreeNode{Val: 4}
		tree.Left.Left = &TreeNode{Val: 11}
		tree.Left.Left.Left = &TreeNode{Val: 7}
		tree.Left.Left.Right = &TreeNode{Val: 2}

		tree.Right = &TreeNode{Val: 8}
		tree.Right.Left = &TreeNode{Val: 13}
		tree.Right.Right = &TreeNode{Val: 4}
		tree.Right.Right.Right = &TreeNode{Val: 1}

		if !hasPathSum(tree, 22) {
			t.Errorf("got: false, want: true")
		}
	})
}
