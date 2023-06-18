package linkedlist

import (
	linkedlist "dsa/DataStructures/LinkedList"
)

func removeNthNode(head *linkedlist.ListNode, n int) *linkedlist.ListNode {
	if head == nil {
		return head
	}
	dummy := &linkedlist.ListNode{Next: head}
	prev, curr := dummy, dummy.Next
	i, l := 1, length(head)
	if l == 1 && n == 1 {
		return nil
	}

	for curr != nil {
		if i == l-n+1 {
			temp := curr
			prev.Next = temp.Next
			break
		}
		i++
		prev = curr
		curr = curr.Next
	}

	return dummy.Next
}

func length(node *linkedlist.ListNode) int {
	length := 0
	for node != nil {
		node = node.Next
		length++
	}

	return length
}

// func main() {
// 	list1 := &linkedlist.ListNode{}
// 	list1 = list1.Build([]int{1})

// 	ans := removeNthNode(list1, 1)
// 	ans.Show()
// }
