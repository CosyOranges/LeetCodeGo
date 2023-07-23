package binarytree

import "testing"

func TestConnect(t *testing.T) {
	test := &Node{Val: 1}
	test.Left = &Node{Val: 2}
	test.Right = &Node{Val: 3}
	test.Right.Right = &Node{Val: 7}
	test.Left.Left = &Node{Val: 4}
	test.Left.Right = &Node{Val: 5}

	ans := &Node{Val: 1}
	ans.Left = &Node{Val: 2}
	ans.Right = &Node{Val: 3}
	ans.Left.Next = ans.Right
	ans.Right.Right = &Node{Val: 7}
	ans.Left.Left = &Node{Val: 4}
	ans.Left.Right = &Node{Val: 5}

	ans.Left.Left.Next = ans.Left.Right
	ans.Left.Right.Next = ans.Right.Right

	t.Run("Test Connect", func(t *testing.T) {
		test = connect(test)

		if !SameTree(test, ans) {
			t.Errorf("Not same tree")
		}
	})
}
