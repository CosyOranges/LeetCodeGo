package binarytree

type Node struct {
	val   int
	Left  *Node
	Right *Node
}

func maxDepth(node *Node) int {
	if node == nil {
		return 0
	}

	depth := 1
	l_depth := maxDepth(node.Left)
	r_depth := maxDepth(node.Right)

	depth += max(l_depth, r_depth)
	return depth
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

// func main() {
// 	root := &Node{val: 3}
// 	root.Left = &Node{val: 9}
// 	root.Right = &Node{val: 20}
// 	root.Right.Left = &Node{val: 15}
// 	root.Right.Right = &Node{val: 7}

// 	fmt.Printf("Max Depth: %v\n", maxDepth(root))
// }
