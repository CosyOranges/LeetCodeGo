package binarytree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	curr  int
	nodes []int
}

func Constructor(root *TreeNode) BSTIterator {
	iter := &BSTIterator{
		curr:  0,
		nodes: []int{},
	}

	var dfs func(iter *BSTIterator, node *TreeNode)

	dfs = func(iter *BSTIterator, node *TreeNode) {
		if node == nil {
			return
		}

		dfs(iter, node.Left)

		iter.nodes = append(iter.nodes, node.Val)

		dfs(iter, node.Right)
	}

	dfs(iter, root)

	return *iter
}

func (this *BSTIterator) Next() int {
	this.curr++
	if this.curr > len(this.nodes) {
		return -10000
	}
	return this.nodes[this.curr-1]
}

func (this *BSTIterator) HasNext() bool {
	if this.curr < len(this.nodes) {
		return true
	}

	return false
}

func main() {
	tree := &TreeNode{Val: 7}
	tree.Left = &TreeNode{Val: 3}
	tree.Right = &TreeNode{Val: 15}
	tree.Right.Left = &TreeNode{Val: 9}
	tree.Right.Right = &TreeNode{Val: 20}

	iter := Constructor(tree)
	fmt.Printf("Iterator: %v\n", iter)
	next := iter.Next()
	fmt.Printf("Next: %v\n", next)
	next = iter.Next()
	fmt.Printf("Next: %v\n", next)
	fmt.Printf("Has Next: %v\n", iter.HasNext())
	next = iter.Next()
	fmt.Printf("Next: %v\n", next)
	fmt.Printf("Has Next: %v\n", iter.HasNext())
	next = iter.Next()
	fmt.Printf("Next: %v\n", next)
	fmt.Printf("Has Next: %v\n", iter.HasNext())
	next = iter.Next()
	fmt.Printf("Next: %v\n", next)
	fmt.Printf("Has Next: %v\n", iter.HasNext())

}
