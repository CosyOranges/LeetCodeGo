package linkedlist

import linkedlist "dsa/DataStructures/LinkedList"

func hasCycle(head *linkedlist.ListNode) bool {
	pos := -1
	if head == nil {
		return false
	}
	posMap := make(map[*linkedlist.ListNode]int)
	pos = 0
	for head.Next != nil {
		if _, ok := posMap[head]; ok {
			return true
		} else {
			posMap[head] = pos
		}
		head = head.Next
		pos++
	}

	return false
}

func hasCycleFloyd(node *linkedlist.ListNode) bool {
	if node == nil {
		return false
	}

	slow, fast := node, node

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

// func main() {
// 	linkedList := &ListNode{}
// 	linkedList = linkedList.Build([]int{3, 2, 0, -4})
// 	linkedList.CreateCycle(2, 4)
// 	fmt.Printf("Has Cycle: %v\n", hasCycleFloyd(linkedList))
// }
