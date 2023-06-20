package binarytree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	t.Run("Test inversion", func(t *testing.T) {
		inverted := &TreeNode{Val: 3}
		inverted.Right = &TreeNode{Val: 9}
		inverted.Left = &TreeNode{Val: 20}
		inverted.Left.Right = &TreeNode{Val: 15}
		inverted.Left.Left = &TreeNode{Val: 7}

		got := InvertTree(root)

		if !reflect.DeepEqual(got, inverted) {
			t.Errorf("Tree was not inverted properly")
			PreOrder(got)
			fmt.Println()
			PreOrder(inverted)
			fmt.Println()
		}
	})
}
