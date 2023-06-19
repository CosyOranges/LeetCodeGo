package linkedlist

import (
	linkedlist "dsa/DataStructures/LinkedList"
)

func rotateRight(head *linkedlist.ListNode, k int) *linkedlist.ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	dummy, curr, breakpoint := &linkedlist.ListNode{Next: head}, head, &linkedlist.ListNode{}
	i, l := 0, length(head)
	point := k % l

	for curr != nil {
		if i == l-(point+1) {
			breakpoint = curr.Next
			curr.Next = nil
		}
		curr = curr.Next
		i++
	}

	dummy.Next = breakpoint
	for breakpoint.Next != nil {
		breakpoint = breakpoint.Next
	}

	breakpoint.Next = head

	return dummy.Next
}

func length(node *linkedlist.ListNode) int {
	length := 0
	for node != nil {
		length++
		node = node.Next
	}
	return length
}

func main() {
	list := &linkedlist.ListNode{}
	list = list.Build([]int{0, 1, 2})
	ans := rotateRight(list, 4)
	ans.Show()
}
