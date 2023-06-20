package binarytree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	root.Left = InvertTree(root.Left)
	root.Right = InvertTree(root.Right)

	return root
}

func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%v ", root.Val)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

// func main() {
// 	root := &TreeNode{Val: 3}
// 	root.Left = &TreeNode{Val: 9}
// 	root.Right = &TreeNode{Val: 20}
// 	root.Right.Left = &TreeNode{Val: 15}
// 	root.Right.Right = &TreeNode{Val: 7}
// 	PreOrder(root)
// 	fmt.Println()
// 	root = InvertTree(root)
// 	PreOrder(root)
// 	fmt.Println()
// }
