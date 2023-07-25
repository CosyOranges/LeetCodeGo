package binarytree

import "testing"

func TestMaxPathSum(t *testing.T) {
	tree := &TreeNode{Val: -10}
	tree.Left = &TreeNode{Val: 9}
	tree.Right = &TreeNode{Val: 20}
	tree.Right.Left = &TreeNode{Val: 15}
	tree.Right.Right = &TreeNode{Val: 7}

	t.Run("Test full tree with negative values", func(t *testing.T) {
		want := 42

		got := maxPathSum(tree)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
