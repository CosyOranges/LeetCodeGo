package linkedlist

import linkedlist "dsa/DataStructures/LinkedList"

func deleteDuplicates(head *linkedlist.ListNode) *linkedlist.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &linkedlist.ListNode{Next: head}
	anchor, curr := dummy, dummy.Next

	for curr != nil {
		if curr.Next != nil && curr.Val == curr.Next.Val {
			for curr.Next != nil && curr.Val == curr.Next.Val {
				curr = curr.Next
			}
			anchor.Next = curr.Next
		} else {
			anchor = anchor.Next
		}
		curr = curr.Next
	}

	return dummy.Next
}

// func main() {
// 	list1 := &linkedlist.ListNode{}
// 	list1 = list1.Build([]int{1, 2, 3, 3})

// 	ans := deleteDuplicates(list1)
// 	ans.Show()
// }
