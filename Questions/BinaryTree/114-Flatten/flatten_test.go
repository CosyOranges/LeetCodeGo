package binarytree

import "testing"

func TestFlatten(t *testing.T) {
	tree := &TreeNode{Val: 1}
	tree.Left = &TreeNode{Val: 2}
	tree.Left.Left = &TreeNode{Val: 3}
	tree.Left.Right = &TreeNode{Val: 4}
	tree.Right = &TreeNode{Val: 5}
	tree.Right.Right = &TreeNode{Val: 6}

	ans := &TreeNode{Val: 1}
	ans.Right = &TreeNode{Val: 2}
	ans.Right.Right = &TreeNode{Val: 3}
	ans.Right.Right.Right = &TreeNode{Val: 4}
	ans.Right.Right.Right.Right = &TreeNode{Val: 5}
	ans.Right.Right.Right.Right.Right = &TreeNode{Val: 6}

	t.Run("Test flatten", func(t *testing.T) {
		flatten(tree)

		if !isSameTree(tree, ans) {
			t.Errorf("Flatten did not work correctly")
		}
	})
}
