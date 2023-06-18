package linkedlist

import (
	linkedlist "dsa/DataStructures/LinkedList"
)

func reverseBetween(head *linkedlist.ListNode, left int, right int) *linkedlist.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// Create a dummy 0th index to our list
	dummy := &linkedlist.ListNode{Next: head}

	// Find our starting index
	// i.e. for 1, 2, 3, 4, 5 this will be 1
	temp := dummy
	for i := 1; i < left; i++ {
		temp = temp.Next
	}

	// To reverse the list we will need 3 extra nodes an "anchor", a "current", and "next"
	// initially these would be: 1, 2, 3
	anchor, curr, next := temp, temp.Next, temp.Next.Next
	for i := left; i < right; i++ {
		// on an iteration shuffled everything along one
		// on the first pass that means: [anchor = 2], [current = 3], [next = 4]
		anchor, curr, next = curr, next, next.Next

		// The reverse happens here where current.Next becomes the anchor
		// e.g. first pass 3.Next = 2
		curr.Next = anchor
	}

	// we stop the above loop with [anchor = 3], [current = 4], [next = 5]
	// Now we need to rearrange the list [temp = 1]
	// 	So temp.Next.Next = next | 1.2.Next = 5
	temp.Next.Next = next
	// And now point temp.Next to current
	//	So 1.Next = 4
	temp.Next = curr
	return dummy.Next
}

// func main() {
// 	list1 := &linkedlist.ListNode{}
// 	list1 = list1.Build([]int{1, 2, 3, 4, 5})
// 	sol := reverseBetween(list1, 2, 4)
// 	sol.Show()
// }
