package binarytree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder, postorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	}

	newNode := &TreeNode{Val: postorder[n-1], Left: nil, Right: nil}
	if len(inorder) == 1 {
		return newNode
	}

	// Find the place to split
	splitIndex := findSplitIndex(inorder, newNode.Val)

	// Build the left subtree
	newNode.Left = buildTree(inorder[:splitIndex], postorder[:len(inorder[:splitIndex])])
	newNode.Right = buildTree(inorder[splitIndex+1:], postorder[len(inorder[:splitIndex]):len(postorder)-1])

	return newNode
}

/*
In: 9,3,15,20,7		Post: 9,15,7,20,3

1st Iteration

	Root: 3

	Build Left:
		In: 9		Post: 9

	Build Right:
		In: 15,20,7	Post: 15,7,20

2nd Iteration

	(Left is finished only handling Right)
	Root: 20

	Build Left:
		In: 15		Post: 15

	Build Right:
		In: 7		Post: 7
*/

func findSplitIndex(inorder []int, val int) int {
	for i, x := range inorder {
		if x == val {
			return i
		}
	}

	return -1
}

func PreOrderArr(node *TreeNode) []int {
	if node == nil {
		return nil
	}

	arr := make([]int, Size(node))

	var i *int
	i = new(int)
	*i = 0
	makePreOrderArr(node, arr, i)

	return arr
}

func makePreOrderArr(node *TreeNode, arr []int, i *int) {
	if node == nil {
		return
	}

	arr[*i] = node.Val
	*i++
	makePreOrderArr(node.Left, arr, i)
	makePreOrderArr(node.Right, arr, i)
}

func Size(node *TreeNode) int {
	if node == nil {
		return 0
	}

	x := 1
	x += Size(node.Left) + Size(node.Right)

	return x
}

func main() {
	in, post := []int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}
	tree := buildTree(in, post)
	preOrder := PreOrderArr(tree)

	fmt.Printf("Pre Order: %v\n", preOrder)
}
