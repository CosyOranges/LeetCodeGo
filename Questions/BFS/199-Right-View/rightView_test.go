package bfs

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	tree := &TreeNode{Val: 1}
	tree.Left = &TreeNode{Val: 2}
	tree.Left.Right = &TreeNode{Val: 5}
	tree.Right = &TreeNode{Val: 3}
	tree.Right.Right = &TreeNode{Val: 4}
	tree.Right.Right.Left = &TreeNode{Val: 20}

	t.Run("Test tree", func(t *testing.T) {
		var want []int = []int{1, 3, 4, 20}

		var got []int = rightSideView(tree)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
