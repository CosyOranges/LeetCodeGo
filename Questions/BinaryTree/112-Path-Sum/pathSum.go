package binarytree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return hasPath(root, 0, targetSum)
}

func hasPath(node *TreeNode, current, target int) bool {
	if node == nil {
		return false
	} else if node.Val+current == target && (node.Left == nil && node.Right == nil) {
		return true
	}

	// Otherwise we still have more options to explore in the left subtree
	if hasPath(node.Left, current+node.Val, target) {
		return true
	}

	// And the right subtree
	if hasPath(node.Right, current+node.Val, target) {
		return true
	}

	return false
}

func main() {
	// tree := &TreeNode{Val: 5}
	// tree.Left = &TreeNode{Val: 4}
	// tree.Left.Left = &TreeNode{Val: 11}
	// tree.Left.Left.Left = &TreeNode{Val: 7}
	// tree.Left.Left.Right = &TreeNode{Val: 2}

	// tree.Right = &TreeNode{Val: 8}
	// tree.Right.Left = &TreeNode{Val: 13}
	// tree.Right.Right = &TreeNode{Val: 4}
	// tree.Right.Right.Right = &TreeNode{Val: 1}

	tree := &TreeNode{Val: -2}
	tree.Right = &TreeNode{Val: -3}

	fmt.Printf("Has path: %v\n", hasPathSum(tree, -5))
}
