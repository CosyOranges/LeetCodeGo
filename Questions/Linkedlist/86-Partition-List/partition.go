package linkedlist

import (
	linkedlist "dsa/DataStructures/LinkedList"
)

func partition(head *linkedlist.ListNode, x int) *linkedlist.ListNode {
	if head == nil {
		return nil
	}

	var tailLeft, headLeft, tailRight, headRight *linkedlist.ListNode
	for head != nil {
		next := head.Next
		// Set head.Next to be nil to prevent creating cycles in new lists
		head.Next = nil

		if head.Val < x {
			if tailLeft == nil && headLeft == nil {
				tailLeft = head
				headLeft = head
			} else {
				tailLeft.Next = head
				tailLeft = head
			}
		} else {
			if tailRight == nil && headRight == nil {
				tailRight = head
				headRight = head
			} else {
				tailRight.Next = head
				tailRight = head
			}
		}

		head = next
	}

	if tailLeft == nil {
		return headRight
	}

	tailLeft.Next = headRight

	return headLeft
}

// func main() {
// 	list := &linkedlist.ListNode{}
// 	list = list.Build([]int{1, 4, 3, 2, 5, 2})

// 	ans := partition(list, 3)
// 	ans.Show()
// }
