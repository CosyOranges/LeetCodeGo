package binarytree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var sum int = 0

	sumNums(root, &sum, 0)

	return sum
}

/*
Node is our current node
Sum is a pointer to the final answer
Temp is the current number down the current path
*/
func sumNums(node *TreeNode, sum *int, temp int) {
	if node == nil {
		return
	}

	temp = temp*10 + node.Val

	// Check if we are at the end of a path
	if node.Left == nil && node.Right == nil {
		*sum += temp
	}

	sumNums(node.Left, sum, temp)
	sumNums(node.Right, sum, temp)

	fmt.Printf("Node: %v | Temp: %v\n", node.Val, temp)
}

func main() {
	tree := &TreeNode{Val: 4}
	tree.Left = &TreeNode{Val: 9}
	tree.Left.Left = &TreeNode{Val: 5}
	tree.Left.Right = &TreeNode{Val: 1}
	tree.Right = &TreeNode{Val: 0}

	fmt.Printf("Ans: %v\n", sumNumbers(tree))
}
