package binarytree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const MIN int = -10000

func maxPathSum(root *TreeNode) int {
	var ans int = MIN
	compute(root, &ans)
	return ans
}

func compute(node *TreeNode, curr *int) int {
	if node == nil {
		return 0
	}

	// compute the left and right total
	left := compute(node.Left, curr)
	right := compute(node.Right, curr)
	fmt.Printf("Left Sum: %v, Right Sum: %v\n", left, right)

	// Now work out whether to take the (left or right) + node.Val or just the node.Val
	maxTemp := max(max(left, right)+node.Val, node.Val)
	fmt.Printf("Max Temp: %v\n", maxTemp)

	// Now determine whether we want to include the root in the current max
	maxParent := max(maxTemp, left+right+node.Val)
	fmt.Printf("Max Parent: %v\n", maxParent)

	*curr = max(maxParent, *curr)

	return maxTemp
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	tree := &TreeNode{Val: -10}
	tree.Left = &TreeNode{Val: 9}
	tree.Right = &TreeNode{Val: 20}
	tree.Right.Left = &TreeNode{Val: 15}
	tree.Right.Right = &TreeNode{Val: 7}

	fmt.Printf("Max Path: %v\n", maxPathSum(tree))
}
