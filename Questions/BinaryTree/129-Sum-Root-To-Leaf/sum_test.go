package binarytree

import "testing"

func TestSumNumbers(t *testing.T) {
	t.Run("Test an empty tree", func(t *testing.T) {
		want := 0

		got := sumNumbers(nil)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test a real tree", func(t *testing.T) {
		want := 1026
		tree := &TreeNode{Val: 4}
		tree.Left = &TreeNode{Val: 9}
		tree.Left.Left = &TreeNode{Val: 5}
		tree.Left.Right = &TreeNode{Val: 1}
		tree.Right = &TreeNode{Val: 0}

		got := sumNumbers(tree)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
