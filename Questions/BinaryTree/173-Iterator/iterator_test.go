package binarytree

import "testing"

func TestIterator(t *testing.T) {
	tree := &TreeNode{Val: 7}
	tree.Left = &TreeNode{Val: 3}
	tree.Right = &TreeNode{Val: 15}
	tree.Right.Left = &TreeNode{Val: 9}
	tree.Right.Right = &TreeNode{Val: 20}

	iterator := Constructor(tree)

	t.Run("Test Construction", func(t *testing.T) {
		if len(iterator.nodes) != 5 {
			t.Errorf("Nodes member is missing elements from tree: %v", iterator.nodes)
		}

		if iterator.curr != 0 {
			t.Errorf("Starting index incorrect, should be 0, is: %v", iterator.curr)
		}
	})

	t.Run("Test Next", func(t *testing.T) {
		want := 3

		got := iterator.Next()

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test HasNext", func(t *testing.T) {
		if !iterator.HasNext() {
			t.Errorf("Should have next")
		}
	})

	t.Run("Test all the way through", func(t *testing.T) {
		iterator.Next()
		iterator.Next()
		iterator.Next()
		iterator.Next()
		iterator.Next()
		if iterator.HasNext() {
			t.Errorf("Should NOT have next")
		}

		test := iterator.Next()

		if test != -10000 {
			t.Errorf("got: %v, want: -10000", test)
		}
	})
}
