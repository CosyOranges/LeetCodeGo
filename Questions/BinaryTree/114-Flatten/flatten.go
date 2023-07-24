package binarytree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	var nodes []*TreeNode
	enqueue(&nodes, root)

	PreOrder(root.Left, &nodes)
	PreOrder(root.Right, &nodes)

	var x *TreeNode
	for i := 0; i < len(nodes); i++ {
		if x != nil {
			x.Right = nodes[i]
			x.Left = nil
		}
		x = nodes[i]
	}

	fmt.Printf("%v\n", root)
}

func PreOrder(node *TreeNode, nodes *[]*TreeNode) {
	if node == nil {
		return
	}

	enqueue(nodes, node)
	PreOrder(node.Left, nodes)
	PreOrder(node.Right, nodes)
}

func enqueue(nodes *[]*TreeNode, item *TreeNode) {
	if nodes == nil {
		return
	}

	*nodes = append(*nodes, item)
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
	tree := &TreeNode{Val: 1}
	tree.Left = &TreeNode{Val: 2}
	tree.Left.Left = &TreeNode{Val: 3}
	tree.Left.Right = &TreeNode{Val: 4}
	tree.Right = &TreeNode{Val: 5}
	tree.Right.Right = &TreeNode{Val: 6}

	flatten(tree)
}
