package binarytree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q != nil {
		return false
	} else if q == nil && p != nil {
		return false
	} else if q == nil && p == nil {
		return true
	}

	if q.Val != p.Val {
		return false
	}

	if !isSameTree(p.Left, q.Left) {
		return false
	}
	if !isSameTree(p.Right, q.Right) {
		return false
	}

	return true
}

func main() {
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

	fmt.Printf("Max Depth: %v\n", isSameTree(root, two))
}
