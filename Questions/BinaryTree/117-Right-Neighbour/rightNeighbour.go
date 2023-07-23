package binarytree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	// First we need a queue to traverse in level order
	var queue []*Node
	enqueue(&queue, root)

	for len(queue) > 0 {
		n := len(queue)

		// fmt.Printf("Node: %v | Left: %v | Right: %v | Next: %v \n", dequedNode.Val, dequedNode.Left, dequedNode.Right, dequedNode.Next)

		var next *Node

		// Now we can loop along the queue and have a trailing "next" variable
		// This trailing node will pick up the current dequeued node, and on the next loop will assign
		// The next dequeued node as the .Next pointer
		// Because it is looping for 'n' it will only ever loop for the width of the current level
		// the "next" var then gets re-assigned to be nil
		for i := 0; i < n; i++ {
			dequedNode := dequeue(&queue)

			if next != nil {
				next.Next = dequedNode
			}

			next = dequedNode

			// Enqueue children
			if dequedNode.Left != nil {
				enqueue(&queue, dequedNode.Left)
			}
			if dequedNode.Right != nil {
				enqueue(&queue, dequedNode.Right)
			}
		}

	}

	return root
}

func enqueue(queue *[]*Node, node *Node) {
	if queue == nil {
		return
	}

	*queue = append(*queue, node)
}

func dequeue(queue *[]*Node) *Node {
	if queue == nil || len(*queue) == 0 {
		return nil
	}

	item := (*queue)[0]
	*queue = (*queue)[1:]

	return item
}

func SameTree(tree1, tree2 *Node) bool {
	if tree1 == nil && tree2 != nil {
		return false
	} else if tree2 == nil && tree1 != nil {
		return false
	} else if tree1 == nil && tree2 == nil {
		return true
	}

	if tree1.Val != tree2.Val {
		return false
	}

	if tree1.Next != nil && tree2.Next != nil {
		if tree1.Next.Val != tree2.Next.Val {
			return false
		}
	} else if tree1.Next != nil && tree2.Next == nil {
		return false
	} else if tree1.Next == nil && tree2.Next != nil {
		return false
	}

	if !SameTree(tree1.Left, tree2.Left) {
		return false
	}

	if !SameTree(tree1.Right, tree2.Right) {
		return false
	}

	return true
}

func main() {
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Right.Right = &Node{Val: 7}
	root.Left.Left = &Node{Val: 4}
	root.Left.Right = &Node{Val: 5}

	root = connect(root)
}
