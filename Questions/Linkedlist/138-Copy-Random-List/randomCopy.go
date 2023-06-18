package linkedlist

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// Create a map to store our nodes and intiate with nils
	nodeMap := make(map[*Node]*Node)
	nodeMap[nil] = nil
	curr := head

	// Loop through list and add all nodes to the map
	for curr != nil {
		copyNode := &Node{Val: curr.Val, Next: nil, Random: nil}
		nodeMap[curr] = copyNode

		curr = curr.Next
	}

	// Reset current to head
	curr = head

	// Connect up all the nodes in the hashmap
	for curr != nil {
		copyNode := nodeMap[curr]

		copyNode.Next = nodeMap[curr.Next]
		copyNode.Random = nodeMap[curr.Random]
		curr = curr.Next
	}

	return nodeMap[head]
}
