package binarytree

import "testing"

func TestIsSameTree(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	two := &TreeNode{Val: 3}
	two.Left = &TreeNode{Val: 9}
	two.Right = &TreeNode{Val: 20}
	two.Right.Left = &TreeNode{Val: 15}
	two.Right.Right = &TreeNode{Val: 9}

	t.Run("Test same tree", func(t *testing.T) {
		if !isSameTree(root, root) {
			t.Errorf("got: false, want: true")
		}
	})

	t.Run("Test diff tree", func(t *testing.T) {
		if isSameTree(root, two) {
			t.Errorf("got: true, want: false")
		}
	})
}
