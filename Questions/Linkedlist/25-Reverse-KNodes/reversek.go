package linkedlist

import (
	linkedlist "dsa/DataStructures/LinkedList"
)

// Algorithm explanation:
// Rejection Cases:
// 	- If head is nil, return head
// 	- If k is 1, return head

// Otherwise:
// 	- Get the length of the list
// 	- Create a 0th index dummy node
// 	- Set prev to 0th index dummy node
// 	- Set curr to head of list

// 	- Loop whilst length is >= k
// 		- Curr = prev.Next
// 		- Loop for i = 1 until i < k (i.e. loop through the KGroup)
// 			- At each index create a temp (next) variable that points to the next node of the current
// 			- Now we have to do some swapping
// 				curr.Next = next.Next
// 				next.Next = prev.Next (which is also current)
// 				prev.Next = next
// 		- prev = current
// 		- reduce length by a factor of k

/*
	Example:
	[1, 2, 3, 4, 5 ,6]

	length = 6
	dummy = 0 | prev = 0 | curr = 1

	1st Iteration:
		curr = 1
		loop i = 1 while i < 3; i++
			[i = 1]
			next = curr.next 		| 2
			curr.Next = next.Next 	| 3
			next.Next = prev.Next	| 1
			prev.Next = next		| 2
			list now looks like:
				[0 -> 2 -> 1 -> 3 -> 4 -> 5 -> 6]
			[i = 2]
			next = curr.next		| 3
			curr.Next = next.Next	| 4
			next.Next = prev.Next	| 2
			prev.Next = next		| 3
			list now looks liks:
				[0 -> 3 -> 2 -> 1 -> 4 -> 5 -> 6]
		length = 6 - 3 = 3
		prev = curr					| 1

	2nd Iteration:
		curr = prev.Next			| 4
		loop i = 1 while i < 3; i++
			[i = 1]
			next = curr.next		| 5
			curr.Next = next.next	| 6
			next.Next = prev.Next	| 4
			prev.Next = next		| 5
			list now looks like:
				[0 -> 3 -> 2 -> 1 -> 5 -> 4 -> 6]
			[i = 2]
			next = curr.next		| 6
			curr.Next = next.Next	| null
			next.Next = prev.Next	| 5
			prev.Next = next		| 6
			list now looks like:
				[0 -> 3 -> 2 -> 1 -> 6 -> 5 -> 4 -> null]
		length = 3 - 3 = 0
		prev = curr					| 4
	l < k BREAK

	return dummy.Next				| 3
*/

func reverseKGroup(head *linkedlist.ListNode, k int) *linkedlist.ListNode {
	if head == nil || k == 1 {
		return head
	}

	l := length(head)
	dummy, curr := &linkedlist.ListNode{Next: head}, head
	prev := dummy

	for l >= k {
		curr = prev.Next
		for i := 1; i < k; i++ {
			next := curr.Next
			curr.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
		}
		prev = curr
		l -= k
	}

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

// func main() {
// 	list1 := &linkedlist.ListNode{}
// 	list1 = list1.Build([]int{1, 2, 3, 4, 5})
// 	sol := reverseKGroup(list1, 2)
// 	sol.Show()
// }
